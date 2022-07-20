import { styled } from "../../theme"
import {
  layout,
  position,
  flexbox,
  typography,
  color,
  display,
  flex,
  space,
  border,
  width,
  fontSize,
} from "styled-system"
import { whiteSpace } from "./custom-style-props"

export const Box = styled.div`
  ${layout}
  ${display}
  ${position}
  ${flexbox}
  ${flex}
  ${space}
  ${color}
  ${typography}
  ${whiteSpace}
  ${fontSize}
  ${display};
  ${width};
  ${space};
  ${border}
  b {
    ${fontSize}
  }
`

Box.defaultProps = {
  display: "flex",
  position: "relative",
}
