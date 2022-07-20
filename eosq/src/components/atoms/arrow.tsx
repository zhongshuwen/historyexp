import * as React from "react"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { styled } from "../../theme"
import { color } from "styled-system"

export const ArrowTo: React.ComponentType<any> = styled(FontAwesomeIcon)`
  height: auto;
  margin: 1px 4px 0 4px;
  vertical-align: middle;
  color: #65656f;
  ${color}
`
