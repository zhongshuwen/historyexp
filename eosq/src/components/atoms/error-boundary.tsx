import React from "react"
import { debugLog } from "../../helpers/log"
import { DataError } from "./data-error"

interface State<E = any> {
  error?: E
}

export interface Props<E = any> {
  errorComponent?: React.ComponentType<{ error: E }>
}

export class ErrorBoundary<E = any> extends React.Component<Props<E>, State<E>> {
  state: State<E> = {}

  public static getDerivedStateFromError(error: Error) {
    return { error }
  }

  public componentDidCatch(error: Error, errorInfo: React.ErrorInfo) {
    debugLog("error boundary catched error %j", errorInfo, error)
  }

  public render() {
    if (this.state.error) {
      const ErrorComponent = this.props.errorComponent || DataError

      return <ErrorComponent error={this.state.error} />
    }

    return this.props.children
  }
}
