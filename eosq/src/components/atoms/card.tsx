import { styled } from "../../theme"
import { position, background, border, boxShadow, padding, borderRadius } from "styled-system"
import { STANDARD_PADDING, STANDARD_SHADOW, MOBILE_PADDING, MEDIA_QUERIES } from "../../theme"

export const Card = styled.div`
  ${position}
  ${background}
  ${border}
  ${boxShadow}
  ${padding}
  ${borderRadius}
  ${MEDIA_QUERIES.smallOnly} {
    padding: ${MOBILE_PADDING};
  }
`

Card.defaultProps = {
  position: "relative",
  background: "white",
  border: "1px solid #eceff1",
  boxShadow: STANDARD_SHADOW,
  padding: STANDARD_PADDING,
  borderRadius: "4px",
  [MEDIA_QUERIES.smallOnly]: {
    padding: MOBILE_PADDING,
  },
}
