import * as React from "react"
import { Cell } from "../ui-grid/ui-grid.component"
import { color } from "styled-system"
import { theme, styled } from "../../theme"

const StyledUl: React.ComponentType<any> = styled.ul`
  list-style: none;
  padding: 0;
`

const StyledLi: React.ComponentType<any> = styled.li`
  list-style: none;

  padding-left: 50px;
  padding-right: 20px;
  padding-bottom: 5px;
  padding-top: 5px;
  ${color};
`

interface LabelValue {
  label: string
  value: string
}

interface Props {
  tabData: LabelValue[]
  onSelectTab: (value: string) => void
}

interface State {
  currentTab: string
}

export class VerticalTabs extends React.Component<Props, State> {
  state: { currentTab: string }

  constructor(props: Props) {
    super(props)
    this.state = { currentTab: this.props.tabData[0].value }
  }

  selectTab(value: string) {
    this.setState({ currentTab: value })
    this.props.onSelectTab(value)
  }

  render() {
    return (
      <Cell bg={"#1C3253"} width="auto">
        <StyledUl>
          {this.props.tabData.map((entry) => {
            return (
              <StyledLi
                key={entry.value}
                bg={this.state.currentTab === entry.value ? "rgba(0, 117, 255, 0.27)" : "#1C3253"}
                color="#fff"
                onClick={() => this.selectTab(entry.value)}
              >
                {entry.label}
              </StyledLi>
            )
          })}
        </StyledUl>
      </Cell>
    )
  }
}
