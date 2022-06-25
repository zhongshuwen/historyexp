import {
  alignItems,
  alignSelf,
  borderColor,
  borderRadius,
  borders,
  color,
  display,
  fontFamily,
  gridColumn,
  gridColumnGap,
  gridRow,
  gridRowGap,
  gridTemplateColumns,
  gridTemplateRows,
  height,
  justifySelf,
  lineHeight,
  minHeight,
  minWidth,
  space,
  textAlign,
  width,
  maxWidth,
  maxHeight,
  fontWeight,
  fontSize,
  compose
} from "styled-system"
import { styled } from "../../theme"
import * as React from "react"

export interface CellExtraProps {
  overflowX?: string
  overflow?: string
  wordBreak?: string
  whiteSpace?: string
  clear?: string
  transition?: string
  float?: string
  cursor?: string
  left?: string
  right?: string
  verticalAlign?: string
}

// FIXME: Complete grid support
const gridStyle = compose(
  alignSelf,
  alignItems,
  justifySelf,
  gridColumnGap,
  gridRowGap,
  gridRow,
  gridColumn,
  gridTemplateColumns,
  gridTemplateRows,
  height,
  width,
  minHeight,
  minWidth,
  space,
  color,
  textAlign,
  borders,
  borderColor,
  borderRadius,
  maxWidth,
  maxHeight
)

export const Grid: React.ComponentType<any> = styled.div`
  display: grid;
  position: relative;
  overflow: ${(props: CellExtraProps) => props.overflow};
  overflow-x: ${(props: CellExtraProps) => props.overflowX};
  
  ${(props) => props.theme.overlays.BG_OVERLAY_TESTNET};
  ${gridStyle}
`

const cellStyle = compose(
  fontSize,
  space,
  color,
  display,
  alignSelf,
  justifySelf,
  gridColumn,
  gridRow,
  height,
  width,
  minHeight,
  minWidth,
  textAlign,
  fontFamily,
  lineHeight,
  borders,
  borderColor,
  borderRadius,
  alignItems,
  maxWidth,
  maxHeight,
  fontWeight
)

export const Cell: React.ComponentType<any> = styled.div`
  position: relative;
  overflow: ${(props: CellExtraProps) => props.overflow};
  overflow-x: ${(props: CellExtraProps) => props.overflowX};
  word-break: ${(props: CellExtraProps) => props.wordBreak};
  white-space: ${(props: CellExtraProps) => props.whiteSpace};
  clear: ${(props: CellExtraProps) => props.clear};
  transition: ${(props: CellExtraProps) => props.transition};
  float: ${(props: CellExtraProps) => props.float};
  :hover {
    cursor: ${(props: CellExtraProps) => props.cursor};
  }
  left: ${(props: CellExtraProps) => props.left};
  right: ${(props: CellExtraProps) => props.right};
  vertical-align: ${(props: CellExtraProps) => props.verticalAlign};
  ${cellStyle}
`

export const HeaderCell: React.ComponentType<any> = styled.div`
  position: relative;
  overflow: ${(props: CellExtraProps) => props.overflow};
  overflow-x: ${(props: CellExtraProps) => props.overflowX};
  word-break: ${(props: CellExtraProps) => props.wordBreak};
  white-space: ${(props: CellExtraProps) => props.whiteSpace};
  clear: ${(props: CellExtraProps) => props.clear};
  transition: ${(props: CellExtraProps) => props.transition};
  float: ${(props: CellExtraProps) => props.float};
  :hover {
    cursor: ${(props: CellExtraProps) => props.cursor};
  }
  left: ${(props: CellExtraProps) => props.left};
  right: ${(props: CellExtraProps) => props.right};
  vertical-align: ${(props: CellExtraProps) => props.verticalAlign};

  
  ${cellStyle}


  position: fixed !important;
  z-index: 1000;
  width: 100%;
  top: 0px;

  background: rgba(80,80,80,0.25);
`

