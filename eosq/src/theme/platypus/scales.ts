// Typographic Scale (numbers are converted to px values)
export const fontSizes = [10, 12, 14, 16, 20, 24, 32, 40, 48, 64, 72]
export const lineHeights = ["18px", "22px", "26px", "34px", "40px", "50px"]

// Spacing Scale (used for margin and padding)
export const space = [0, 4, 8, 16, 32, 64, 128, 256, 512]

export const breakPoints = {
  small: 768,
  medium: 1280,
  large: 1440
}

export const mediaQueries = {
  smallOnly: `@media (max-width: ${breakPoints.small - 1}px)`,
  small: `@media (min-width: ${breakPoints.small}px)`,
  medium: `@media (min-width: ${breakPoints.medium}px)`,
  large: `@media (min-width: ${breakPoints.large}px)`
}


export const STANDARD_PADDING = "20px 50px"
export const MOBILE_PADDING = "20px 12px"
export const STANDARD_SHADOW = "0 2px 4px rgba(0, 0, 0, .05), 0 2px 10px rgba(0, 0, 0, .05)"
export const RADIUS = "4px"

export const BREAKPOINTS = {
  small: 768,
  medium: 1280,
  large: 1440,
}

export const MEDIA_QUERIES = {
  smallOnly: `@media (max-width: ${BREAKPOINTS.small - 1}px)`,
  small: `@media (min-width: ${BREAKPOINTS.small}px)`,
  medium: `@media (min-width: ${BREAKPOINTS.medium}px)`,
  large: `@media (min-width: ${BREAKPOINTS.large}px)`,
}
