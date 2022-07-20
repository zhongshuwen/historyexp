import { styled } from "../../theme"
import { color } from "styled-system"

export const SiderMenu = styled.div<{ alignCenter?: boolean }>`
  display: flex;
  flex-direction: row;
  max-width: 100%;
  align-items: center;
  white-space: nowrap;
  margin-left: auto;
  ${color}
  font-family: "Roboto", sans-serif;
  font-weight: 400;
  font-size: 12px;

  ${(props) => (props.alignCenter === true ? "align-items: center;" : "")}
  a {
    margin-left: 30px;
    ${color}
    :hover {
      color: "white";
    }
  }
`

SiderMenu.defaultProps = {
  color: "#bbc7d3",
}
