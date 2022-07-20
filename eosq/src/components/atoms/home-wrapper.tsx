import * as React from "react"
import { styled } from "../../theme"
import { color } from "styled-system"
import { Box } from "./box"

export const HomeWrapper: React.ComponentType<any> = styled(Box)`
  min-height: 100vh;
  ${color}
`
