import { styled } from "../../theme"
import { display, layout, flexbox } from "styled-system"

export const Sider = styled.div`
  ${display}
  ${flexbox}
  ${layout}
`

Sider.defaultProps = {
  display: "flex",
  flexDirection: "row",
  alignItems: "center",
  justifyContent: "center",
  margin: "0px",
}
