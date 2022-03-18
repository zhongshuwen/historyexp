import { t } from "i18next"
import { observer } from "mobx-react"
import * as React from "react"
import { RouteComponentProps } from "react-router-dom"
import { ServerError } from "../../components/server-error/server-error.component"
import { Cell, Grid } from "../../atoms/ui-grid/ui-grid.component"
import { debugLog } from "../../services/logger"
import { metricsStore } from "../../stores"
import { ContentLoaderComponent } from "../../components/content-loader/content-loader.component"
import { Panel } from "../../atoms/panel/panel.component"
import { TransactionDetailHeader } from "./summary/transaction-detail-header"
import { PanelTitleBanner } from "../../atoms/panel/panel-title-banner"
import { Text } from "../../atoms/text/text.component"
import { WrappingText } from "../../atoms/text-elements/misc"
import { DataLoading, DataError, BULLET, truncateString } from "@dfuse/explorer"
import { PageContainer } from "../../components/page-container/page-container"
import { TransactionLifecycle, Stream } from "@dfuse/client"
import { TransactionContents } from "./transaction-contents"
import { computeTransactionTrustPercentage } from "../../models/transaction"
import { registerTransactionLifecycleListener } from "../../streams/transaction-listeners"
import { TransactionLifecycleWrap } from "../../services/transaction-lifecycle"

export interface PathParams {
  id: string
}

interface Props extends RouteComponentProps<PathParams> {}

interface State {
  fetchError: boolean
  lifecycleWrap: TransactionLifecycleWrap | undefined
}

@observer
export class TransactionDetailPage extends ContentLoaderComponent<Props, State> {
  handlerMetricsId?: string
  statusUpdated = false
  state: State = { lifecycleWrap: undefined, fetchError: false }
  transactionStream?: Stream

  componentDidMount = async () => {
    await this.registerStreams()

    this.changeDocumentTitle()
  }

  get trustPercentage() {
    if (this.state.lifecycleWrap) {
      return computeTransactionTrustPercentage(
        this.state.lifecycleWrap.blockNum,
        metricsStore.headBlockNum,
        metricsStore.lastIrreversibleBlockNum
      )
    }

    return 0
  }

  async registerStreams() {
    this.setState({ fetchError: false })
    this.transactionStream = await registerTransactionLifecycleListener(
      this.props.match.params.id,
      (lifecycle: TransactionLifecycle) => {
        if (lifecycle && lifecycle !== null) {
          this.setState({ lifecycleWrap: new TransactionLifecycleWrap(lifecycle) })
        }
      },
      () => {
        this.setState({ fetchError: true })
      }
    )
  }

  async unregisterStreams() {
    if (this.transactionStream) {
      await this.transactionStream.close()
      this.transactionStream = undefined
    }
  }

  componentDidUpdate(prevProps: Props) {
    if (prevProps.match.params.id !== this.props.match.params.id) {
      this.changeDocumentTitle()
      this.unregisterStreams()
      this.registerStreams()
    }
  }

  changeDocumentTitle() {
    document.title = `${truncateString(this.props.match.params.id, 8).join(
      ""
    )} ${BULLET} 交易 ${BULLET} 国家级版权交易保护联盟链区块浏览器`
  }

  componentWillUnmount() {
    this.unregisterStreams()
  }

  renderError = (error?: any) => {
    debugLog("Handling transaction stream error.", error)
    return <ServerError />
  }

  renderLoading = (message: string) => {
    return (
      <Grid px={[4]} py={[2]}>
        <DataLoading text={message} />
      </Grid>
    )
  }

  renderNotSeenYet(transactionId: string) {
    return [
      <PanelTitleBanner key="0" title={t("transaction.banner.title")} content={transactionId} />,
      <Grid key="1" gridRowGap={[3]}>
        <Panel>
          <Cell p={[3, 4]}>
            <Text>{t("transaction.notSeenYet.notFound")}</Text>
            <Text>{t("transaction.notSeenYet.watchingForNetwork")}</Text>&nbsp;
            <WrappingText fontWeight="bold" color="secondHighlight">
              {transactionId}
            </WrappingText>
          </Cell>
        </Panel>
      </Grid>
    ]
  }

  renderContent = () => {
    if (!this.state.lifecycleWrap) {
      return this.renderLoading(this.props.match.params.id)
    }

    if (this.state.fetchError) {
      return (
        <PageContainer>
          <DataError />
        </PageContainer>
      )
    }

    return (
      <PageContainer>
        <PanelTitleBanner
          title={t("transaction.banner.title")}
          content={this.props.match.params.id+"afawef"}
        />
        <Grid gridRowGap={[3]}>
          <Panel>
            <TransactionDetailHeader lifecycleWrap={this.state.lifecycleWrap} />
          </Panel>
          <TransactionContents
            history={this.props.history}
            location={this.props.location}
            currentTab="actions"
            lifecycleWrap={this.state.lifecycleWrap}
          />
        </Grid>
      </PageContainer>
    )
  }

  render() {
    return this.renderContent()
  }
}
