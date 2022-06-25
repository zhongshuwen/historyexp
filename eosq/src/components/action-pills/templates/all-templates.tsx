import { PillComponentClass } from "./generic-pill.component"
import { BuyRamPillComponent } from "./system/buy-ram-pill.component"
import { BuyRamBytesPillComponent } from "./system/buy-ram-bytes-pill.component"
import { RefundPillComponent } from "./system/refund-pill.component"
import { DelegateBandwidthPillComponent } from "./system/delegate-bandwidth-pill.component"
import { IssuePillComponent } from "./system/issue-pill.component"
import { TransferPillComponent } from "./transfer-pill.component"
import { VotePillComponent } from "./system/vote-pill.component"
import { UnDelegateBandwidthPillComponent } from "./system/undelegate-bandwidth-pill.component"
import { NewAccountPillComponent } from "./system/newaccount-pill.component"
import { LinkAuthPillComponent } from "./system/linkauth-pill.component"
import { UpdateAuthPillComponent } from "./system/updateauth-pill.component"
import { ClaimRewardsPillComponent } from "./system/claim-rewards-pill.component"
import { SetcodePillComponent } from "./system/setcode-pill.component"
import { RegProxyPillComponent } from "./system/regproxy-pill.component"
import { DfuseEventPillComponent } from "./dfuse-events/dfuse-event-pill.component"
export const ALL_TEMPLATES: PillComponentClass[] = [
  BuyRamPillComponent,
  BuyRamBytesPillComponent,
  RefundPillComponent,
  DelegateBandwidthPillComponent,
  IssuePillComponent,
  TransferPillComponent,
  VotePillComponent,
  UnDelegateBandwidthPillComponent,
  NewAccountPillComponent,
  LinkAuthPillComponent,
  UpdateAuthPillComponent,
  ClaimRewardsPillComponent,
  SetcodePillComponent,
  RegProxyPillComponent,
  DfuseEventPillComponent,

]
