import emotionStyled, { CreateStyled } from "@emotion/styled"
import { injectThemedStyled } from "@dfuse/explorer"
import { colors } from "./colors"
import { breakPoints, mediaQueries, fontSizes, lineHeights, space } from "./scales"
import { Scale, get as resolveValueWithDefault } from "styled-system"
import * as overlays from './core/overlays';
export const theme = {
  breakPoints,
  mediaQueries,
  fontSizes,
  lineHeights,
  space,
  colors,
  overlays,
  fontFamily: {
    roboto: "OPPOSans2, sans-serif",
    mono: "'Roboto Mono', monospace",
    robotoCondensed: "'OPPOSans2', sans-serif",
    opensans: "Open Sans",
    iceland: "Iceland",
    lato: "OPPOSans2"
  },
}

export function resolveValue(n: number | string, scale?: Scale) {
  return resolveValueWithDefault(scale, n, n)
}

export type ThemeInterface = typeof theme

export const styled = emotionStyled as CreateStyled<ThemeInterface>

injectThemedStyled(styled)
