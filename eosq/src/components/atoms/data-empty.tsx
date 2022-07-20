import * as React from "react"
import * as CSS from "csstype"
import { SubTitle } from "./typography"
import { Box } from "./box"

type Props = {
  text: string
  color?: CSS.ColorProperty
}

export const DataEmpty: React.FC<Props> = ({ text, color, children }) => (
  <Box px={[4]} py={[4]} align="center" justify="center" flexDirection="column" w={["100%"]}>
    <SubTitle color={color || "#6a74de"} fontWeight={["normal"]} fontSize={[3]}>
      {text}
    </SubTitle>
    {children}
  </Box>
)
