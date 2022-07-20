import { styled } from "../../theme"
import { color, layout, display, position, grid, flexbox, typography, space } from "styled-system"
import { cursor, float, clear } from "./custom-style-props"

export const Cell = styled.div`
  ${display}
  ${position}
  ${color}
  ${layout}
  ${grid}
  ${flexbox}
  ${typography}
  ${cursor}
  ${float}
  ${clear}
  ${space}
`

Cell.defaultProps = {
  position: "relative",
}
