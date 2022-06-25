import { t } from "i18next"
import * as React from "react"
import { ExternalTextLink, Text } from "../../atoms/text/text.component"
import { Cell, Grid } from "../../atoms/ui-grid/ui-grid.component"
import { translate } from "react-i18next"
import { Links } from "../../routes"
import { Link } from "react-router-dom"
import { fontSize, space } from "styled-system"
import { theme, styled } from "../../theme"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { BULLET, NBSP } from "@dfuse/explorer"
import { getCurrentLanguageValue } from "../settings-selectors/settings.helpers"

const LogoLink: React.ComponentType<any> = styled(Link)`
  display: block;
  display: flex;
  align-items: left;
  justify-content: left;
  ${space};
`

const LogoFirst: React.ComponentType<any> = styled.div`
font-family: ${(props) => props.theme.fontFamily.lato};
  font-weight: 600;
  color: ${(props) => props.theme.colors.logo1};
  ${fontSize};
  top: -10px;
  position: relative;
`

const LogoSecond: React.ComponentType<any> = styled.div`
  font-family: ${(props) => props.theme.fontFamily.lato};
  font-weight: 600;
  color: ${(props) => props.theme.colors.logo2};
  ${fontSize};
  top: -10px;
  position: relative;
`

const StyledText: React.ComponentType<any> = styled(Text)`
  &:hover {
    color: ${theme.colors.primary};
  }
`

const StyledFont: React.ComponentType<any> = styled(FontAwesomeIcon)`
  &:hover {
    color: ${theme.colors.primary};
  }
`

const BaseFooter = () => (
  <Cell height="auto" bg={"#f8f8f8"} mt="10px" pt="6px">
    <Cell
      maxWidth={["1800px"]}
      px={[4]}
      mx="auto"
      lineHeight={[2]}
      pt={[1]}
      pb={[2]}
      my={[4]}
      width="100%"
      textAlign="center"
    >
      <a
        href={`https://zhongshuwen.com/${getCurrentLanguageValue()}`}
        target="_blank"
        rel="noopener noreferrer"
      >
      <img src="/images/logo-explorer-black.png" alt="联盟链区块浏览器" style={{width:"30vw",maxWidth:"300px"}}/>
      </a>
    </Cell>
    <Cell p={[3]} mt={[3]} borderTop={`1px solid ${theme.colors.bleu10}`} textAlign="center">
      <Text color={theme.colors.primary} fontSize={[1]}>
        {t("footer.copyright", { year: new Date().getFullYear() })}
      </Text>
      <Text color={theme.colors.primary} fontSize={[1]}>
        <ExternalTextLink fontSize={[1]} to={t("footer.privacyPolicyLink")}>
          <StyledText display="inline-block" fontSize={[1]} color={theme.colors.bleu6}>
            {t("footer.privacyPolicy")}
          </StyledText>
        </ExternalTextLink>
        {NBSP}
        {BULLET}
        {NBSP}
        <ExternalTextLink fontSize={[1]} to={t("footer.termsOfServicesLink")}>
          <StyledText display="inline-block" fontSize={[1]} color={theme.colors.bleu6}>
            {t("footer.termsOfServices")}
          </StyledText>
        </ExternalTextLink>
      </Text>
    </Cell>
  </Cell>
)

export const Footer = translate()(BaseFooter)
