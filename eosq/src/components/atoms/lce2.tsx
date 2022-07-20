import * as React from "react"
import { DataLoading } from "./data-loading"
import { ErrorBoundary } from "./error-boundary"

type Props = {
  loadingText?: string
  loadingComponent?: React.ComponentType<{ text?: string }>
  errorComponent?: React.ComponentType<{ error: any }>
  children?: React.ReactNode
}

/**
 * A Load Content Error (LCE) base component for React.Suspense usage.
 *
 * TBC...
 */
export const LCE2: React.FC<Props> = ({
  loadingText,
  loadingComponent,
  errorComponent,
  children,
}) => {
  const LoadingComponent = loadingComponent || DataLoading

  return (
    <ErrorBoundary errorComponent={errorComponent}>
      <React.Suspense fallback={<LoadingComponent text={loadingText} />}>{children}</React.Suspense>
    </ErrorBoundary>
  )
}
