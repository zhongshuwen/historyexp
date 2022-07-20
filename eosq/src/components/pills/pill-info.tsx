import * as React from "react"
import { Cell } from "../atoms/cell"
import { ellipsis } from "../../helpers/ellipsis"
import { Text } from "../atoms/typography"

import { styled } from "../../theme"

const Container = styled(Cell)`
  align-self: center;
  border-left: 1px dotted #aaa;
`

const InfoText = styled(Text)`
  align-self: center;
  font-family: "'Roboto Condensed', sans-serif";
  font-size: 14px;
  padding: 0 16px 0 8px;
  color: black;
`

interface Props {
  info: string
}

export const PillInfo: React.FC<Props> = ({ info }) => (
  <Container>
    <InfoText>{ellipsis(info, 50)}</InfoText>
  </Container>
)
