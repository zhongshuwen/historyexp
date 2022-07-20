import * as React from "react"
import { Grid } from "../atoms/grid"
import { Text } from "../atoms/typography"
import { styled, MEDIA_QUERIES } from "../../theme"

interface Props {
  label: string
  label2?: string
  variant?: "standard" | "compact" | "auto"
  mb?: string | number
  ml?: string | number
  color?: string
}

const Container = styled(Grid)<{ gridTemplateColumns: string[] }>`
  grid-template-columns: ${({ gridTemplateColumns }) => gridTemplateColumns[0]};
  align-items: center;
  grid-gap: 0px 12px;

  ${MEDIA_QUERIES.medium} {
    grid-template-columns: ${({ gridTemplateColumns }) => gridTemplateColumns[1]};
  }
`

export const CellValue = styled.div`
  position: relative;
  word-break: break-word;
  white-space: normal;
`

export const CellLabel = styled.div`
  text-align: center;
  min-height: 40px;
`

export const DetailLine: React.FC<Props> = ({
  label: label1,
  label2,
  color,
  variant,
  ml,
  children,
}) => {
  let templateColumns = ["1fr", "2fr 3fr"]
  if (variant === "compact") {
    templateColumns = ["minmax(70px, 1fr) minmax(400px, 6fr)"]
  } else if (variant === "auto") {
    templateColumns = ["1fr", "auto 3fr"]
  }

  let labelTemplateRows = ["1fr"]
  if (label2) {
    labelTemplateRows = ["1fr 1fr"]
  }

  return (
    <Container gridTemplateColumns={templateColumns}>
      <Grid
        gridTemplateColumns="auto"
        gridTemplateRows={labelTemplateRows}
        height="100%"
        alignItems="start"
        gridGap="10px 0px"
      >
        <Text fontWeight="bold" lineHeight="23px" color={color}>
          {label1}&nbsp;
        </Text>
        {label2 && (
          <Text fontWeight="bold" lineHeight="23px" color={color}>
            {label2}&nbsp;
          </Text>
        )}
      </Grid>
      <CellValue>
        <Grid gridTemplateColumns="auto" gridGap="10px 0px" ml={ml}>
          {children}
        </Grid>
      </CellValue>
    </Container>
  )
}

export const DetailLineAuto: React.FC<Props> = ({ label, color, mb, children }) => {
  const templateColumns = ["1fr", "auto 3fr"]

  return (
    <Grid mb={[mb !== undefined ? mb : 2]} gridTemplateColumns={templateColumns}>
      <Text color={color || "text"} fontWeight="bold">
        {label}&nbsp;
      </Text>
      <CellValue>{children}</CellValue>
    </Grid>
  )
}
