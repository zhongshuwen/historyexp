import * as CSS from "csstype"
import * as React from "react"
import { styled } from "../../theme"
import { Link } from "react-router-dom"
import {
  alignSelf,
  display,
  fontFamily,
  fontSize as _fontSize,
  fontWeight,
  justifySelf,
  lineHeight,
  textAlign as _textAlign,
  borders,
  width,
  color,
  layout,
  position,
  space,
  typography,
  flex,
  flexbox,
} from "styled-system"

// TODO: Merge text components

/* ETHQ */

const baseStyles = `
  word-break: break-word;
  color: inherit;
`

export interface TextStyleProps {
  size?: "small" | "normal" | "large"
  variant?: "condensed" | "monospace"
  hoverable?: boolean

  bg?: string
  color?: string
  p?: string
  px?: string
  py?: string
  pl?: string
  pr?: string

  m?: string
  mx?: string
  my?: string
  ml?: string
  mr?: string

  position?: CSS.PositionProperty
  display?: CSS.DisplayProperty
  lineHeight?: string
  fontSize?: string
  fontWeight?: number | string
  whiteSpace?: CSS.WhiteSpaceProperty

  width?: string
  height?: string
  minHeight?: string
  minWidth?: string
  maxHeight?: string
  maxWidth?: string

  flex?: string

  // Make those more restrictive by finding a way to remove the generic `string` type
  alignItems?: CSS.AlignItemsProperty
  alignSelf?: CSS.AlignSelfProperty
  justifyContent?: CSS.JustifyContentProperty

  hover?: string
}

export const textStyle = (props: TextStyleProps & React.HTMLAttributes<HTMLElement>) => {
  const lines: string[] = []
  const color = props.color
  let fontSize: string | undefined
  if (props.size === "small") {
    fontSize = "0.8em"
  } else if (props.size === "large") {
    fontSize = "1.2em"
  }

  if (props.variant === "monospace") {
    lines.push('font-family: "Roboto Mono", monospace;')
  } else if (props.variant === "condensed") {
    lines.push('font-family: "Roboto Condensed";')
  }

  if (props.onClick != null || props.hover || props.hoverable) {
    lines.push(`
        &:hover {
          cursor: pointer;
          ${props.hover ? props.hover : ""}
        }
      `)
  }

  if (props.fontSize) {
    fontSize = props.fontSize
  }

  if (props.bg) {
    lines.push(`background-color: ${props.bg};`)
  }

  lines.push(`color: ${color};`)

  if (fontSize) {
    lines.push(`font-size: ${fontSize};`)
  }

  if (props.fontWeight) {
    lines.push(`font-weight: ${props.fontWeight};`)
  }

  if (props.p) {
    lines.push(`padding: ${props.p};`)
  }

  if (props.px) {
    lines.push(`padding-left: ${props.px};`)
    lines.push(`padding-right: ${props.px};`)
  }

  if (props.py) {
    lines.push(`padding-top: ${props.py};`)
    lines.push(`padding-bottom: ${props.py};`)
  }

  if (props.pl) {
    lines.push(`padding-left: ${props.pl};`)
  }

  if (props.pr) {
    lines.push(`padding-right: ${props.pr};`)
  }

  if (props.m) {
    lines.push(`margin: ${props.m};`)
  }

  if (props.mx) {
    lines.push(`margin-left: ${props.mx};`)
    lines.push(`margin-right: ${props.mx};`)
  }

  if (props.my) {
    lines.push(`margin-top: ${props.my};`)
    lines.push(`margin-bottom: ${props.my};`)
  }

  if (props.ml) {
    lines.push(`margin-left: ${props.ml};`)
  }

  if (props.mr) {
    lines.push(`margin-right: ${props.mr};`)
  }

  if (props.display) {
    lines.push(`display: ${props.display};`)
  }

  if (props.position) {
    lines.push(`position: ${props.position};`)
  }

  if (props.lineHeight) {
    lines.push(`line-height: ${props.lineHeight};`)
  }

  if (props.width) {
    lines.push(`width: ${props.width};`)
  }

  if (props.height) {
    lines.push(`height: ${props.height};`)
  }

  if (props.minHeight) {
    lines.push(`min-height: ${props.minHeight};`)
  }

  if (props.minWidth) {
    lines.push(`min-width: ${props.minWidth};`)
  }

  if (props.maxHeight) {
    lines.push(`max-height: ${props.maxHeight};`)
  }

  if (props.maxWidth) {
    lines.push(`max-width: ${props.maxWidth};`)
  }

  if (props.whiteSpace) {
    lines.push(`whitespace: ${props.whiteSpace};`)
  }

  if (props.alignItems) {
    lines.push(`align-items: ${props.alignItems};`)
  }

  if (props.alignSelf) {
    lines.push(`align-self: ${props.alignSelf};`)
  }

  if (props.justifyContent) {
    lines.push(`justify-content: ${props.justifyContent};`)
  }

  return lines.join("\n")
}

export const H1 = styled.h1<TextStyleProps>`
  ${baseStyles}
  ${(props) => textStyle({ ...props, fontSize: "2.25em" })}
`

export const H2 = styled.h2<TextStyleProps>`
  ${baseStyles}
  ${(props) => textStyle({ ...props, fontSize: "1.75em" })}
`

export const H3 = styled.h3<TextStyleProps>`
  ${baseStyles}
  &.underline {
    padding-bottom: 10px;
    border-bottom: 1px solid hsl(370, 5%, 95%);
  }

  ${(props) => textStyle({ ...props, fontSize: "1.5em" })}
`

/* EOSQ */
// TODO: implement hoverable prop and replace all custom textStyle
export const Text: React.ComponentType<any> = styled.div`
  ${flex}
  ${flexbox}
  ${typography}
  ${space}
  ${position}
  ${layout}
  ${display};
  ${_fontSize};
  ${space};
  ${color};
  ${fontWeight};
  ${_textAlign};
  ${fontFamily};
  ${alignSelf};
  ${justifySelf};
  ${lineHeight};
  ${borders};
  ${width};
  text-overflow: ${(props: any) => props.textOverflow};
  text-transform: ${(props: any) => props.textTransform};
  white-space: ${(props: any) => props.whiteSpace};
  word-break: ${(props: any) => props.wordBreak};
`

Text.defaultProps = {
  position: "relative",
}

export type TextProps = React.ComponentProps<typeof Text>

export const HoverableText: React.ComponentType<any> = styled(Text)`
  &:hover {
    cursor: pointer;
    ${color}
  }
`

export const HoverableTextNoHighlight: React.ComponentType<any> = styled(Text)`
  &:hover {
    cursor: pointer;
  }
`

export const Ellipsis: React.ComponentType<any> = styled(Text)`
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
`

export const CondensedBold: React.ComponentType<any> = styled.b`
  font-family: "Roboto Condensed", sans-serif;
  font-weight: 800;
`

export const BigTitle: React.ComponentType<any> = styled.h1`
  ${_fontSize};
  ${space};
  ${color};
  ${fontWeight};
  ${_textAlign};
  ${fontFamily};
  ${alignSelf};
  ${justifySelf};
`

export const Title: React.ComponentType<any> = styled.h2`
  ${_fontSize};
  ${space};
  ${color};
  ${fontWeight};
  ${_textAlign};
  ${fontFamily};
  ${alignSelf};
  ${justifySelf};
`

export const SubTitle: React.ComponentType<any> = styled.h3`
  ${display};
  ${_fontSize};
  ${space};
  ${color};
  ${fontWeight};
  ${_textAlign};
  ${fontFamily};
  ${alignSelf};
  ${justifySelf};
`

Text.defaultProps = {
  color: "text",
}

BigTitle.defaultProps = {
  color: "text",
}

Title.defaultProps = {
  color: "text",
}

SubTitle.defaultProps = {
  color: "text",
  my: [2],
}

export interface TextLinkProps {
  whiteSpace?: string
  lineHeight?: string
  download?: string
  to: string
  fontSize?: any
  fontFamily?: any
  fontWeight?: any
  style?: any
  pt?: any
  pb?: any
  pr?: any
  p?: any
  textAlign?: any
  color?: any
  pl?: any
  width?: any
  mr?: any
  ml?: any
  my?: any
  mx?: any
}

export const LinkStyledText: React.ComponentType<any> = styled(HoverableText)`
  display: inline;
  ${_textAlign};
  ${space};
  ${color};
  ${fontWeight};
  ${_fontSize};
  ${fontFamily};
  ${alignSelf};
  ${justifySelf};
  ${lineHeight};
  ${borders};
  ${width};
`

export const StyledLink: React.ComponentType<any> = styled(Link)`
  ${_fontSize};
`

export const TextLinkLight: React.FC<TextLinkProps> = ({ to, children, ...rest }) => {
  return (
    <Link to={to}>
      <LinkStyledText color="link2" {...rest}>
        {children}
      </LinkStyledText>
    </Link>
  )
}

export const TextLink: React.FC<TextLinkProps> = ({ to, children, ...rest }) => {
  return (
    <StyledLink fontSize={rest && rest.fontSize ? rest.fontSize : ""} to={to}>
      <LinkStyledText color="link" {...rest}>
        {children}
      </LinkStyledText>
    </StyledLink>
  )
}

export const ExternalTextLink: React.FC<TextLinkProps> = ({ to, download, children, ...rest }) => {
  if (download) {
    return (
      <a href={to} target="_blank" rel="noreferrer" download={download}>
        <LinkStyledText color="link" {...rest}>
          {children}
        </LinkStyledText>
      </a>
    )
  }
  return (
    <a href={to} target="_blank" rel="noreferrer" {...download}>
      <LinkStyledText color="link" {...rest}>
        {children}
      </LinkStyledText>
    </a>
  )
}

export const ExternalTextLinkLight: React.FC<TextLinkProps> = ({ to, children, ...rest }) => {
  return (
    <a href={to} target="_blank" rel="noreferrer">
      <LinkStyledText color="link2" {...rest}>
        {children}
      </LinkStyledText>
    </a>
  )
}

export class KeyValueFormatEllipsis extends React.Component<{ content: string }> {
  render() {
    const regex: RegExp = /(\S*: )/g
    return (
      <Ellipsis fontFamily="Roboto Condensed" fontSize={[1]}>
        {this.props.content.split(regex).map((value: string, index: number) => {
          if (regex.test(value)) {
            return <CondensedBold key={index}>{value}</CondensedBold>
          }

          return value
        })}
      </Ellipsis>
    )
  }
}

export const MonospaceTextLink: React.FC<TextLinkProps> = ({ ...rest }) => (
  <TextLink {...rest} fontFamily="'Roboto Mono', monospace;" lineHeight="1em" />
)

export const MonospaceText: React.ComponentType<any> = styled(Text)`
  font-family: "Roboto Mono", monospace;
  white-space: nowrap;
`

export const MonospaceTextWrap: React.ComponentType<any> = styled(Text)`
  font-family: "Roboto Mono", monospace;
`

export const WrappingText: React.ComponentType<any> = styled(Text)`
  overflow: hidden;
  word-wrap: break-word;
`

export const WrappingMonospaceText: React.ComponentType<any> = styled(Text)`
  font-family: "Roboto Mono", monospace;
  overflow: hidden;
  word-wrap: break-word;
`
