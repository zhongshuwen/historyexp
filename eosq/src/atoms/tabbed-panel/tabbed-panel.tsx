import * as React from "react"
import { createStyles, WithStyles, withStyles } from "@material-ui/core/styles"
import Tabs from "@material-ui/core/Tabs"
import Tab from "@material-ui/core/Tab"
import { Cell } from "../ui-grid/ui-grid.component"
import { styled, theme } from "../../theme"
import { FontSizeProps } from "styled-system"
import { Text } from "../text/text.component"

const Wrapper: React.ComponentType<any> = styled(Cell)`
  border: 1px solid ${(props) => props.theme.colors.border};
  grid-auto-flow: row;
  min-width: 0px;
  min-height: 200px;
`

const styles = () =>
  createStyles({
    root: {
      fontSize: "30px !important",
      border: `1px solid ${theme.colors.border}`,
      background: theme.colors.tpRootBackgroundColor,
    },
    tabsRoot: {
      background: theme.colors.tpTabsRootBackgroundColor,
      height: "48px"
    },
    tabsIndicator: {
      display: "none"
    },
    wrapper: {
      alignItems: "flex-start !important",
      paddingLeft: "10px",
      paddingTop: "10px",
      "&:hover": {
        cursor: "auto !important"
      }
    },
    tabRoot: {
      maxWidth: "none",
      "&:last-child": {
        borderRight: "none !important"
      },

      borderRight: `1px solid ${theme.colors.border}`,
      borderBottom: `1px solid ${theme.colors.border}`,
      textTransform: "initial",
      minWidth: 72,
      fontWeight: "normal",
      fontFamily: theme.fontFamily.roboto,
      "&:hover": {
        color: theme.colors.tpTabRootHoverColor,
        opacity: 1
      },
      "&$tabSelected": {
        color: theme.colors.tpTabRootSelectedColor,
        fontWeight: 500,
        borderBottom: "none !important"
      },
      "&:focus": {
        color:theme.colors.tpTabRootFocusColor,
      },
      color: theme.colors.tpTabRootColor
    },
    selected: {
      border: "none !important"
    },
    tabSelected: {
      background:  theme.colors.tpTabSelectedBackground,
      borderRight: `1px solid ${theme.colors.border}`,
      borderBottom: theme.colors.tpTabSelectedBorderBottom,
    }
  })

interface LabelValue {
  label: string
  value: string
}

interface Props extends FontSizeProps {
  tabData: LabelValue[]
  onSelect: (key: string) => void
  selected?: string
}

const decorate = withStyles(styles)

export const TabbedPanel = decorate<any>(
  class extends React.Component<Props & WithStyles<typeof styles>, { value: number }> {
    constructor(props: Props & WithStyles<typeof styles>) {
      super(props)
      this.state = { value: this.getIndexSelected() }
    }

    getIndexSelected = () => {
      const index = this.props.tabData.findIndex((entry) => entry.label === this.props.selected)
      return index === -1 ? 0 : index
    }

    componentDidUpdate(prevProps: Props): void {
      if (prevProps.tabData.length !== this.props.tabData.length && this.props.selected) {
        // eslint-disable-next-line react/no-did-update-set-state
        this.setState({ value: this.getIndexSelected() })
      }
    }

    handleChange = (_: React.ChangeEvent<{}>, value: number) => {
      this.setState({ value })
    }

    renderLabel = (item: LabelValue): JSX.Element => {
      return (
        <Text color="inherit" fontWeight="bold" fontSize={this.props.fontSize || "16px"}>
          {item.value}
        </Text>
      )
    }

    renderTab = (item: LabelValue, index: number) => {
      const { classes } = this.props

      const appliedClasses: any = { root: classes.tabRoot, selected: classes.tabSelected }
      if (this.props.tabData.length <= 1) {
        appliedClasses.wrapper = classes.wrapper
      }

      return (
        <Tab
          key={index}
          onClick={() => this.props.onSelect(item.label)}
          disableRipple={true}
          classes={classes ? appliedClasses : {}}
          label={this.renderLabel(item)}
          style={{ width: "100%", textAlign: "left" }}
        />
      )
    }

    renderTabs() {
      return this.props.tabData.map((item, index) => this.renderTab(item, index))
    }

    render() {
      const { classes } = this.props
      const { value } = this.state

      return (
        <Wrapper className={classes ? classes.root : ""}>
          <Tabs
            value={value}
            onChange={this.handleChange}
            variant="fullWidth"
            classes={classes ? { root: classes.tabsRoot, indicator: classes.tabsIndicator } : {}}
          >
            {this.renderTabs()}
          </Tabs>
          <Cell mt={[3]}>{this.props.children}</Cell>
        </Wrapper>
      )
    }
  }
)
