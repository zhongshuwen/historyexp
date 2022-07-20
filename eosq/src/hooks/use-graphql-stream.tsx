import { useEffect, useState } from "react"
import { GraphqlResponseError, Stream } from "@dfuse/client"
import { DocumentNode } from "graphql"
import { print as printGraphqlDocument } from "graphql/language/printer"
import { getDfuseClient } from "../clients/dfuse"
import {
  PromiseState,
  promiseStatePending,
  promiseStateRejected,
  promiseStateResolved,
} from "./use-promise"

export function useGraphqlStream<T = any>(params: {
  document: string | DocumentNode
  variables: Record<string, unknown>
  onData: (results: T) => void
  onError: (errors: Error[]) => void
  onComplete?: () => void
}): PromiseState<Stream, GraphqlResponseError[]> {
  const { document, variables, onData, onComplete, onError } = params

  const [stream, setStream] = useState<PromiseState<Stream, GraphqlResponseError[]>>(
    promiseStatePending()
  )

  let stringDocument = document as string
  if (typeof document !== "string") {
    stringDocument = printGraphqlDocument(document)
  }

  useEffect(() => {
    setStream(promiseStatePending())
    ;(async () => {
      try {
        const stream = await getDfuseClient().graphql<T>(
          stringDocument,
          (message) => {
            switch (message.type) {
              case "data":
                onData(message.data)
                break
              case "complete":
                onComplete && onComplete()
                break
              case "error":
                onError(message.errors)
                break
            }
          },
          {
            variables,
          }
        )
        setStream(promiseStateResolved(stream))
      } catch (error) {
        setStream(
          promiseStateRejected([
            {
              message: `${error}`,
              extensions: { cause: error },
            },
          ])
        )
      }
    })()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [stringDocument, JSON.stringify(variables)])

  return stream
}
