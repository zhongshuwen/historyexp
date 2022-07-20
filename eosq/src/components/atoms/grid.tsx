import { styled } from "../../theme"
import { color, layout, display, position, grid, flexbox } from "styled-system"
import { cursor } from "./custom-style-props"

export const Grid = styled.div`
  ${display}
  ${position}
    ${color}
    ${layout}
    ${grid}
    ${flexbox}
    ${cursor}
`

Grid.defaultProps = {
  display: "grid",
  position: "relative",
}
