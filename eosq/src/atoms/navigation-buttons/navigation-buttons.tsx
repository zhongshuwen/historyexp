import * as React from "react"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faAngleLeft, faAngleDoubleLeft, faAngleRight } from "@fortawesome/free-solid-svg-icons"
import {FiChevronRight,FiChevronLeft,FiChevronsLeft,FiChevronsRight} from 'react-icons/fi'
import { Cell } from "../ui-grid/ui-grid.component"
import { borders, color as color_ } from "styled-system"
import { theme, styled } from "../../theme"

interface NavigationProps {
  direction: "next" | "previous" | "first"
  onClick?: () => void
  variant?: string
}

const ColoredTile: React.ComponentType<any> = styled(Cell)`
  ${borders}
  ${color_}

  width: 32px;
  height: 32px;
  text-align: center;
  cursor: pointer;
  display: flex;
  justify-items: center;
  align-items: center;
  justify-content: center;
  background:transparent !important;
  opacity:0.9;



background: linear-gradient(126.97deg, rgba(6, 11, 38, 0.74) 28.26%, rgba(26, 31, 55, 0.5) 91.2%) !important;
border-radius: 10px;
border:1px solid #304060 !important;

& svg {
  opacity:0.7;
}
  &:hover {
    opacity:1;
    background-color: ${(props) => props.hoverBg};
    color: ${(props) => (props.hoverColor||"rgba(255,255,255,0.8)")};
    &> * {
      opacity:1;
    }
  }
  &:active {
    background: linear-gradient(126.97deg, rgba(6, 11, 58, 0.94) 28.26%, rgba(26, 31, 85, 0.7) 91.2%) !important;

  }
`

export const NavigationButton: React.SFC<NavigationProps> = ({ direction, onClick, variant }) => {
  variant = variant || "dark"
  const color = variant === "dark" ? "#fff" : "#fff"
  const hoverColor = variant === "dark" ? "#fff" : "#fff"
  
  const hoverBg = variant === "dark" ? theme.colors.bleu10 : theme.colors.ternary
  const size = "1.5em"
  const ChevronLeft = <FiChevronLeft size={size} />
  const ChevronRight = <FiChevronRight size={size} />
  
  //<FontAwesomeIcon size="3x" color={color} icon={FiChevronLeft} />
  const DoubleChevronLeft = <FiChevronsLeft size={size} />

  let Chevron = ChevronLeft

  if (direction === "next") {
    Chevron = ChevronRight
  } else if (direction === "first") {
    Chevron = DoubleChevronLeft
  }

  return (
    <ColoredTile
      hoverColor={hoverColor}
      color={color}
      hoverBg={hoverBg}
      bg={variant === "light" ? "#fff" : theme.colors.bleu9}
      border={variant === "light" ? `1px solid ${theme.colors.grey2}` : "none"}
      title={direction}
      size="48"
      onClick={onClick}
    >
      {Chevron}
    </ColoredTile>
  )
}

interface NavigationButtonsProps {
  onNext: () => void
  onPrev: () => void
  onFirst?: () => void
  showFirst: boolean
  showNext: boolean
  showPrev: boolean
  variant?: string
  textAlign?: string
}

export class NavigationButtons extends React.Component<NavigationButtonsProps> {
  renderNavigationButton(direction: "next" | "previous" | "first", display: boolean) {
    let onClick: () => void
    if (direction === "next") {
      onClick = this.props.onNext
    } else if (direction === "previous") {
      onClick = this.props.onPrev
    } else if (direction === "first" && this.props.onFirst) {
      onClick = this.props.onFirst
    }

    if (display) {
      return (
        <NavigationButton
          variant={this.props.variant}
          direction={direction}
          onClick={() => onClick()}
        />
      )
    }

    return null
  }
  render() {
    return (
      <Cell textAlign={this.props.textAlign} display="flex" alignItems="center" justifyContent="center" justifyItems="center">
        {this.props.onFirst ? (
          <Cell display="inline-block" p={1}>
            {this.renderNavigationButton("first", this.props.showFirst)}
          </Cell>
        ) : null}
        <Cell display="inline-block" p={1}>
          {this.renderNavigationButton("previous", this.props.showPrev)}
        </Cell>
        <Cell display="inline-block" p={1}>
          {this.renderNavigationButton("next", this.props.showNext)}
        </Cell>
      </Cell>
    )
  }
}
