import * as React from "react"
import { Action } from "@dfuse/client"
import { shallowWithTheme } from "../../../../tests/renderers"
import { getActionMock } from "../../../../__mocks__/transaction.mock"
import { BuyRamBytesPillComponent } from "../system/buy-ram-bytes-pill.component"
import { BuyRamPillComponent } from "../system/buy-ram-pill.component"
import { ClaimRewardsPillComponent } from "../system/claim-rewards-pill.component"
import { DelegateBandwidthPillComponent } from "../system/delegate-bandwidth-pill.component"
import { IssuePillComponent } from "../system/issue-pill.component"
import { LinkAuthPillComponent } from "../system/linkauth-pill.component"
import { NewAccountPillComponent } from "../system/newaccount-pill.component"
import { RefundPillComponent } from "../system/refund-pill.component"
import { RegProxyPillComponent } from "../system/regproxy-pill.component"
import { TransferPillComponent } from "../transfer-pill.component"
import { UnDelegateBandwidthPillComponent } from "../system/undelegate-bandwidth-pill.component"
import { UpdateAuthPillComponent } from "../system/updateauth-pill.component"
import { VotePillComponent } from "../system/vote-pill.component"
import { GenericPillComponent } from "../generic-pill.component"
import { SetcodePillComponent } from "../system/setcode-pill.component"

const headerAndTitleOptions = {
  header: {
    color: "#33333",
    text: "text",
    hoverTitle: "hovetTitle"
  },
  title: "title"
}

describe("Templates", () => {
  describe("BuyRamBytesPillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          payer: "payer",
          receiver: "receiver",
          bytes: 150
        }
      })

      expect(shallowWithTheme(renderComponent(BuyRamBytesPillComponent, action))).toMatchSnapshot()
    })
  })

  describe("BuyRamPillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          payer: "payer",
          receiver: "receiver",
          quantity: "20.0000 EOS"
        }
      })

      expect(shallowWithTheme(renderComponent(BuyRamPillComponent, action))).toMatchSnapshot()
    })
  })

  describe("ClaimRewardsPillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          owner: "owner"
        }
      })

      expect(shallowWithTheme(renderComponent(ClaimRewardsPillComponent, action))).toMatchSnapshot()
    })
  })


  describe("DelegateBandwidthPillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          from: "from",
          receiver: "receiver",
          stake_cpu_quantity: "20.0000 EOS",
          stake_net_quantity: "10.0000 EOS"
        }
      })

      expect(
        shallowWithTheme(renderComponent(DelegateBandwidthPillComponent, action))
      ).toMatchSnapshot()
    })
  })

  describe("GenericPillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          random_field: "random_field",
          other_field: "other_field"
        }
      })

      expect(shallowWithTheme(renderComponent(GenericPillComponent, action))).toMatchSnapshot()
    })

    it("should render properly when no data but hex data present", () => {
      const action = getActionMock({
        data: undefined
      })

      expect(shallowWithTheme(renderComponent(GenericPillComponent, action))).toMatchSnapshot()
    })

    it("should render properly when no data and no hex data present", () => {
      const action = getActionMock({
        data: undefined
      })

      delete action.hex_data

      expect(shallowWithTheme(renderComponent(GenericPillComponent, action))).toMatchSnapshot()
    })
  })

  describe("IssuePillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          to: "eoscanadacom",
          quantity: "10.0000 EOS"
        }
      })

      expect(shallowWithTheme(renderComponent(IssuePillComponent, action))).toMatchSnapshot()
    })
  })

  describe("LinkAuthPillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          account: "eoscanadacom",
          requirement: "requirement",
          type: "type"
        }
      })

      expect(shallowWithTheme(renderComponent(LinkAuthPillComponent, action))).toMatchSnapshot()
    })
  })

  describe("NewAccountPillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          creator: "eoscanadacom",
          name: "newaccount",
          owner: "newaccount",
          active: true
        }
      })

      expect(shallowWithTheme(renderComponent(NewAccountPillComponent, action))).toMatchSnapshot()
    })
  })

  describe("RefundPillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          owner: "newaccount"
        }
      })

      expect(shallowWithTheme(renderComponent(RefundPillComponent, action))).toMatchSnapshot()
    })
  })

  describe("RegProxyPillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          isproxy: false,
          proxy: "testing"
        }
      })

      expect(shallowWithTheme(renderComponent(RegProxyPillComponent, action))).toMatchSnapshot()
    })
  })


  describe("BaseSetcodePillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          code: "code",
          account: "account"
        }
      })

      expect(shallowWithTheme(renderComponent(SetcodePillComponent, action))).toMatchSnapshot()
    })
  })

  describe("TransferPillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          from: "from",
          to: "to",
          quantity: "2.000 EOS"
        }
      })

      expect(shallowWithTheme(renderComponent(TransferPillComponent, action))).toMatchSnapshot()
    })
  })

  describe("UnDelegateBandwidthPillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          from: "from",
          unstake_net_quantity: "1.000 EOS",
          unstake_cpu_quantity: "2.000 EOS"
        }
      })

      expect(
        shallowWithTheme(renderComponent(UnDelegateBandwidthPillComponent, action))
      ).toMatchSnapshot()
    })
  })

  describe("UpdateAuthPillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          auth: "auth",
          parent: "parent",
          permission: "permission"
        }
      })

      expect(shallowWithTheme(renderComponent(UpdateAuthPillComponent, action))).toMatchSnapshot()
    })
  })

  describe("VotePillComponent", () => {
    it("should render properly", () => {
      const action = getActionMock({
        data: {
          voter: "voter",
          producers: ["producer"]
        }
      })

      expect(shallowWithTheme(renderComponent(VotePillComponent, action))).toMatchSnapshot()
    })
  })
})

function renderComponent(Component: any, action: Action<any>) {
  return <Component action={action} headerAndTitleOptions={headerAndTitleOptions} />
}
