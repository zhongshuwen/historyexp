import { t } from "i18next"
import { observer } from "mobx-react"
import * as React from "react"
import { formatAmount, formatNumber, Box } from "../../dexplorer"
import { styled, theme } from "../../theme"
import { metricsStore } from "../../stores"
import { Text } from "../../atoms/text/text.component"

import { Cell, Grid } from "../../atoms/ui-grid/ui-grid.component"
import { BannerContainer, BannerItem, BannerTitle } from "./banner-item.component"
import { AmountVariation } from "./variation.component"
import { Link } from "react-router-dom"
import { Links } from "../../routes"
import { Config } from "../../models/config"

const BannerWrapper: React.ComponentType<any> = styled(Grid)`
  grid-column-gap: 28px;
  border-style: solid;
  border-color: ${(props) => props.theme.colors.bleu6};
`

const Price: React.ComponentType<any> = styled(Text)`
  color: ${(props) => props.theme.colors.bannerValue};
  line-height: 1;
  font-weight: 700;
  font-family: ${(props) => props.theme.fontFamily.robotoCondensed};
`

const BannerMarketPrice: React.SFC<{ price: number; variation: number }> = ({
  price,
  variation
}) => {
  let formattedPrice = formatAmount(price)
  if (price < 0) {
    formattedPrice = ""
  }

  return (
    <BannerContainer>
      <Box flexDirection="row">
        <Box align={["center", "center", "left"]}>
          <Price fontSize={[8]}>{formattedPrice}</Price>
          <Box
            align="left"
            ml={[2]}
            flexDirection="column"
            justifyContent="center"
            alignItems="center"
          >
            <Cell>
              <BannerTitle fontSize={[1]}>{t("banner.eos_usd")}</BannerTitle>
              <AmountVariation variation={variation} textColor="bannerValue" />
            </Cell>
          </Box>
        </Box>
      </Box>
    </BannerContainer>
  )
}

@observer
export class Banner extends React.Component {
  renderProducerLink(account: string) {
    if (!account || account.length === 0) {
      return (
        <Text
          fontFamily={theme.fontFamily.robotoCondensed+";"}
          fontWeight="bold"
          color={theme.colors.bannerTextColor}
          fontSize={[4, 5, 5]}
        >
          {account}
        </Text>
      )
    }
    return (
      <Link to={Links.viewAccount({ id: account })}>
        <Text
          fontFamily={theme.fontFamily.robotoCondensed+";"}
          fontWeight="bold"
          color={theme.colors.bannerTextColor}
          fontSize={[4, 5, 5]}
        >
          {account}
        </Text>
      </Link>
    )
  }

  renderBlockLink(blockNum: string, blockId: string) {
    if (!blockId || blockId.length === 0) {
      return (
        <Text
          fontFamily={theme.fontFamily.robotoCondensed+";"}
          fontWeight="bold"
          color={theme.colors.bannerTextColor}
          fontSize={[4, 5, 5]}
        >
          {blockNum}
        </Text>
      )
    }
    return (
      <Link to={Links.viewBlock({ id: blockId })}>
        <Text
          fontFamily={theme.fontFamily.robotoCondensed+";"}
          fontWeight="bold"
          color={theme.colors.bannerTextColor}
          fontSize={[4, 5, 5]}
        >
          {blockNum}
        </Text>
      </Link>
    )
  }


  render() {
    return (
      <Cell>
        <BannerWrapper
          borderLeft={["0px", "1px"]}
          borderRight={["0px"]}
          borderBottom={["0px"]}
          borderTop={["0px"]}
          py="0px"
          gridTemplateColumns={["3fr 2fr", "1fr 3fr 3fr 3fr"]}
        >
          <div></div>
          <BannerItem
            
            title={t("banner.head_block")}
            details={formatNumber(metricsStore.headBlockNum)}
          />
          <BannerItem
            title={t("banner.irreversible_block")}
            titleTip={t("banner.irreversible_block_tooltip")}
            details={this.renderBlockLink(
              formatNumber(metricsStore.lastIrreversibleBlockNum),
              metricsStore.lastIrreversibleBlockId
            )}
          />
          <BannerItem
            title={t("banner.head_block_producer")}
            titleTip={t("banner.head_block_producer_tooltip")}
            details={this.renderProducerLink(metricsStore.headBlockProducer)}
          />
        </BannerWrapper>
      </Cell>
    )
  }
}
