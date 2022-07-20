import * as React from "react"
import { styled } from "../../theme"
import { Text } from "./typography"
import { Box } from "./box"
import { LoadingOutlined } from "@ant-design/icons"

const Loader = styled(LoadingOutlined)`
  font-size: 36px;
  width: 100%;
`

type Props = {
  text?: string
  color?: string
}

const renderText = (text?: string, color?: string) => {
  if (text === undefined) {
    return null
  }

  return (
    <Text
      wordBreak="break-all"
      whiteSapce="normal"
      py={[3]}
      fontSize={[3]}
      color={[color || "text"]}
    >
      {text}
    </Text>
  )
}

export const DataLoading: React.FC<Props> = ({ text, color, children }) => (
  <Box
    pt={[4]}
    pb={[4]}
    textAlign="center"
    justify="center"
    flexDirection="column"
    width={["100%"]}
  >
    <Loader color={color || "text"} />
    {renderText(text, color)}
    {children}
  </Box>
)
