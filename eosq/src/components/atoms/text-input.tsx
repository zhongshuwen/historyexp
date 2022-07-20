import { styled } from "../../theme"
import { color, border, boxShadow, layout, typography } from "styled-system"
import { STANDARD_SHADOW, MEDIA_QUERIES } from "../../theme"

export const TextInput = styled.input`
  ${boxShadow}
  ${border}
  ${typography}
  ${color}
  ${layout}
`

TextInput.defaultProps = {
  boxShadow: STANDARD_SHADOW,
  border: "none",
  fontWeight: 500,
  color: "hsl(370, 0%, 45%)",
  height: "30px",
  borderRadius: "4px",
  width: "600px",
  padding: "30px 20px",
  paddingRight: 0,
  fontSize: "18px",
  [MEDIA_QUERIES.smallOnly]: {
    fontSize: "14px",
    padding: "25px 15px",
    paddingRight: 0,
  },
}
