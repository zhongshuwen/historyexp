import * as React from "react"
import { Box } from "../atoms/box"
import { Cell } from "../atoms/cell"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faMinus } from "@fortawesome/free-solid-svg-icons/faMinus"
import { faPlus } from "@fortawesome/free-solid-svg-icons/faPlus"
import Tooltip from "antd/es/tooltip"
import {
  PillWrapper,
  PillContainer,
  PillContainerDetails,
  PillExpandedContainer,
  PillHeaderText,
  PillInfoContainer,
  PillOverviewRow,
  PillClickable,
  AnimatedPillContainer,
  PillLogoContainer,
  PillLogo,
  PillWithRigthInfo,
} from "./pill-elements"

interface Props {
  headerHoverTitle?: string
  info?: string
  disabled?: boolean
  failed?: boolean
  failureMessage?: string
  highlighted?: boolean

  // Present or not, at the very beginning of the pill
  logo?: PillLogoProps

  // First section of the pill, we call it the header
  headerText: JSX.Element | string
  headerColor?: string
  headerBgColor?: string
  headerHoverText?: string

  // Second section of the pill, we call it the title, can be empty to render a 2px box
  isSubCall?: boolean
  percentage?: number
  selfGas?: string | number
  cumulativeGas?: string | number

  // Leftmost colored "title" part of pill
  pill2: JSX.Element[] | JSX.Element | null

  // Background of expand button
  expandButtonBgColor?: string
  expandButtonColor?: string

  renderInfo: () => JSX.Element[] | JSX.Element | null

  content: JSX.Element[] | JSX.Element | string | null
  renderExpandedContent: () => JSX.Element[] | JSX.Element | null
}

export interface PillLogoProps {
  path: string
  website: string
}

interface State {
  isOpen: boolean
}

export class Pill extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props)
    this.state = {
      isOpen: false,
    }
  }

  openWebsiteLink = () => {
    window.open(this.props.logo!.website, "_blank")
  }

  toggleIsOpen = () => {
    if (this.props.disabled) {
      return
    }

    this.setState((prevState: State) => ({ isOpen: !prevState.isOpen }))
  }

  renderOverviewRow() {
    return (
      <PillOverviewRow
        bg={this.props.failed ? "#FF91A0" : this.props.highlighted ? "lightyellow" : "white"}
        minHeight="26px"
      >
        {this.props.pill2}
        {this.props.content}
        {this.props.disabled ? null : (
          <PillClickable
            bg={this.props.expandButtonBgColor || "grey"}
            px="12px"
            alignItems="center"
          >
            <FontAwesomeIcon
              size="sm"
              color={this.props.expandButtonColor || "white"}
              icon={this.state.isOpen ? faMinus : faPlus}
            />
          </PillClickable>
        )}
      </PillOverviewRow>
    )
  }

  renderHeader(text: JSX.Element | string, color: string, title: string) {
    const WrapperComponent = this.props.disabled ? Box : PillClickable

    return (
      <WrapperComponent bg={color} alignItems="center" justifyContent="center">
        <PillHeaderText
          title={title}
          pl={this.props.logo ? "35px" : "16px"}
          pr="15px"
          color={this.props.headerColor || "white"}
          fontSize="12px"
        >
          {text}
        </PillHeaderText>
      </WrapperComponent>
    )
  }

  renderLogo() {
    if (this.props.logo) {
      return (
        <PillLogoContainer>
          <PillLogo onClick={this.openWebsiteLink}>
            <img width="100%" alt={this.props.logo.website} src={this.props.logo.path} />
          </PillLogo>
        </PillLogoContainer>
      )
    }

    return null
  }

  render() {
    const { isOpen } = this.state
    const infoSection = this.props.renderInfo()

    let infoPadding = "15px"
    if (infoSection == null) {
      infoPadding = "0"
    }
    return (
      <PillWithRigthInfo gridTemplateColumns={"1fr"}>
        <div />
        <PillWrapper width="100%" display="block" clear="both" my="5px">
          {this.renderLogo()}
          <PillContainer
            cursor={this.props.disabled ? "default" : "pointer"}
            failed={this.props.failed}
            overflow="hidden"
            onClick={this.props.disabled ? () => {} : this.toggleIsOpen}
            gridTemplateColumns="auto 1fr"
          >
            {this.props.failed && (
              <Tooltip
                placement="left"
                mouseEnterDelay={0.01}
                mouseLeaveDelay={0.15}
                title={this.props.failureMessage || "Unknown error"}
              >
                {this.renderHeader(
                  this.props.headerText,
                  this.props.failed ? "#FF91A0" : this.props.headerBgColor || "grey",
                  this.props.headerHoverText || ""
                )}
              </Tooltip>
            )}
            {!this.props.failed &&
              this.renderHeader(
                this.props.headerText,
                this.props.failed ? "#FF91A0" : this.props.headerBgColor || "grey",
                this.props.headerHoverText || ""
              )}
            {this.renderOverviewRow()}
          </PillContainer>
          <AnimatedPillContainer
            pl="31px"
            pr="35px"
            isOpen={isOpen}
            maxHeight={isOpen ? "3000px" : "0px"}
          >
            <PillContainerDetails>
              {/* The element right under the expanded pill, right before the actual expanded content */}
              <PillInfoContainer withBorderBottom={infoSection != null}>
                <Cell p={infoPadding}>{infoSection}</Cell>
              </PillInfoContainer>

              {/* The actual expanded content, usually tab based */}
              <PillExpandedContainer bg="white">
                {this.props.renderExpandedContent()}
              </PillExpandedContainer>
            </PillContainerDetails>
          </AnimatedPillContainer>
        </PillWrapper>
      </PillWithRigthInfo>
    )
  }
}
