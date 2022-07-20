import { styled, MEDIA_QUERIES } from "../../theme"

export const SmallOnly = styled.div`
  ${MEDIA_QUERIES.small} {
    display: none;
  }
`

export const Small = styled.div`
  display: none;

  ${MEDIA_QUERIES.small} {
    display: inherit;
  }
`

export const Medium = styled.div`
  display: none;

  ${MEDIA_QUERIES.medium} {
    display: inherit;
  }
`

export const Large = styled.div`
  display: none;

  ${MEDIA_QUERIES.large} {
    display: inherit;
  }
`
