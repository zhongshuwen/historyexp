import { styled, MEDIA_QUERIES } from "../../theme"

export const SpaceBetween = styled.div`
  display: flex;
  flex-direction: row;
  width: 100%;
  justify-content: space-between;
  align-items: center;

  & > * {
    display: flex;
    flex: 0 1 auto;
    max-width: 100%;
  }

  ${MEDIA_QUERIES.smallOnly} {
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;

    & > * {
      width: 100%;
    }
  }
`
