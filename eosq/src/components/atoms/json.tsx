import * as React from "react"
import { styled } from "../../theme"
import { display, layout, typography, fontFamily, fontSize } from "styled-system"
import { whiteSpace, wordWrap } from "./custom-style-props"

const JsonContainerCode = styled.code`
  ${display}
  ${layout}
${typography}
${fontFamily}
${whiteSpace}
${wordWrap}
`

JsonContainerCode.defaultProps = {
  fontFamily: "Roboto Mono, monospace",
  whiteSpace: "pre-wrap",
  wordWrap: "break-word",
  display: "block",
  overflowX: "hidden",
}

const JsonContainerPre = styled.pre`
  ${fontSize}
  ${display}
${typography}
${whiteSpace}
`

JsonContainerPre.defaultProps = {
  fontSize: "15px",
  whiteSpace: "pre-wrap",
  display: "block",
  overflowX: "hidden",
}

export const JsonWrapper: React.FC<any> = ({ fontSize, children }) => (
  <JsonContainerPre>
    <JsonContainerCode fontSize={fontSize}>{children}</JsonContainerCode>
  </JsonContainerPre>
)
