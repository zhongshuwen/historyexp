/**
 * This is an inline-copy of https://github.com/andreiduca/use-async-resource.
 *
 * Mainly copied to remove cache behavior and easier fixing (if needed). Should think
 * about using the library directly at some point.
 */
import { useCallback, useMemo, useRef, useState } from "react"

/**
 * A typical api function: takes an arbitrary number of arguments of type A
 * and returns a Promise which resolves with a specific response type of R.
 */
export type ApiFn<R, A extends any[] = []> = (...args: A) => Promise<R>

/**
 * An updater function: has a similar signature with the original api function,
 * but doesn't return anything because it only triggers new api calls.
 */
export type UpdaterFn<A extends any[] = []> = (...args: A) => void

/**
 * A simple data reader function: returns the response type R.
 */
type DataFn<R> = () => R
/**
 * A lazy data reader function: can return the response type R or `undefined`.
 */
type LazyDataFn<R> = () => R | undefined

/**
 * A modifier function which takes as only argument the response type R and returns a different type M.
 */
export type ModifierFn<R, M = any> = (response: R) => M

/**
 * A data reader with a modifier function,
 * returning the modified type M instead of the response type R.
 */
type ModifiedDataFn<R> = <M>(modifier: ModifierFn<R, M>) => M
/**
 * A lazy data reader with a modifier function,
 * returning the modified type M instead of the response type R, or `undefined`.
 */
type LazyModifiedDataFn<R> = <M>(modifier: ModifierFn<R, M>) => M | undefined

// Finally, our actual eager and lazy implementations will use both versions (with and without a modifier function),
// so we need overloaded types that will satisfy them simultaneously

/**
 * A data reader function with an optional modifier function.
 */
export type DataOrModifiedFn<R> = DataFn<R> & ModifiedDataFn<R>
/**
 * A lazy data reader function with an optional modifier function.
 */
export type LazyDataOrModifiedFn<R> = LazyDataFn<R> & LazyModifiedDataFn<R>

/**
 * Wrapper for an apiFunction without params.
 * It only takes the api function as an argument.
 * It returns a data reader with an optional modifier function.
 *
 * @param apiFn A typical api function that doesn't take any parameters.
 */
export function initializeDataReader<ResponseType>(
  apiFn: ApiFn<ResponseType>
): DataOrModifiedFn<ResponseType>

/**
 * Wrapper for an apiFunction with params.
 * It takes the api function and all its expected arguments.
 * Also returns a data reader with an optional modifier function.
 *
 * @param apiFn A typical api function with parameters.
 * @param parameters An arbitrary number of parameters.
 */
export function initializeDataReader<ResponseType, ArgTypes extends any[]>(
  apiFn: ApiFn<ResponseType, ArgTypes>,
  ...parameters: ArgTypes
): DataOrModifiedFn<ResponseType>

// implementation that covers the above overloads
export function initializeDataReader<ResponseType, ArgTypes extends any[] = []>(
  apiFn: ApiFn<ResponseType, ArgTypes>,
  ...parameters: ArgTypes
) {
  type AsyncStatus = "init" | "done" | "error"

  let data: ResponseType
  let status: AsyncStatus = "init"
  let error: any

  const fetchingPromise = apiFn(...parameters)
    .then((result) => {
      data = result
      status = "done"
    })
    .catch((err) => {
      error = err
      status = "error"
    })

  // the return type successfully satisfies DataOrModifiedFn<ResponseType>
  function dataReaderFn(): ResponseType
  function dataReaderFn<M>(modifier: ModifierFn<ResponseType, M>): M
  function dataReaderFn<M>(modifier?: ModifierFn<ResponseType, M>) {
    if (status === "init") {
      throw fetchingPromise
    } else if (status === "error") {
      throw error
    }

    return typeof modifier === "function" ? (modifier(data) as M) : (data as ResponseType)
  }

  return dataReaderFn
}

/**
 * Lazy initializer.
 * The only param passed is the api function that will be wrapped.
 * The returned data reader LazyDataOrModifiedFn<ResponseType> is "lazy",
 *   meaning it can return `undefined` if the api call hasn't started.
 * The returned updater function UpdaterFn<ArgTypes>
 *   can take any number of arguments, just like the wrapped api function
 *
 * @param apiFunction A typical api function.
 */
export function useAsyncResource<ResponseType, ArgTypes extends any[]>(
  apiFunction: ApiFn<ResponseType, ArgTypes>
): [LazyDataOrModifiedFn<ResponseType>, UpdaterFn<ArgTypes>]

/**
 * Eager initializer for an api function without params.
 * The second param must be `[]` to indicate we want to start the api call immediately.
 * The returned data reader DataOrModifiedFn<ResponseType> is "eager",
 *   meaning it will always return the ResponseType
 *   (or a modified version of it, if requested).
 * The returned updater function doesn't take any arguments,
 *   just like the wrapped api function
 *
 * @param apiFunction A typical api function that doesn't take any parameters.
 * @param eagerLoading If present, the api function will get executed immediately.
 */
export function useAsyncResource<ResponseType>(
  apiFunction: ApiFn<ResponseType>,
  eagerLoading: never[] // the type of an empty array `[]` is `never[]`
): [DataOrModifiedFn<ResponseType>, UpdaterFn]

/**
 * Eager initializer for an api function with params.
 * The returned data reader is "eager", meaning it will return the ResponseType
 *   (or a modified version of it, if requested).
 * The returned updater function can take any number of arguments,
 *   just like the wrapped api function
 *
 * @param apiFunction A typical api function with an arbitrary number of parameters.
 * @param parameters If present, the api function will get executed immediately with these parameters.
 */
export function useAsyncResource<ResponseType, ArgTypes extends any[]>(
  apiFunction: ApiFn<ResponseType, ArgTypes>,
  ...parameters: ArgTypes
): [DataOrModifiedFn<ResponseType>, UpdaterFn<ArgTypes>]

// implementation that covers the above overloads
export function useAsyncResource<ResponseType, ArgTypes extends any[]>(
  apiFunction: ApiFn<ResponseType> | ApiFn<ResponseType, ArgTypes>,
  ...parameters: ArgTypes
) {
  // keep the data reader inside a mutable object ref
  // always initialize with a lazy data reader, as it can be overwritten by the useMemo immediately
  const dataReaderObj = useRef<DataOrModifiedFn<ResponseType> | LazyDataOrModifiedFn<ResponseType>>(
    () => undefined
  )

  // like useEffect, but runs immediately
  useMemo(() => {
    if (parameters.length) {
      // eager initialization for api functions that don't accept arguments
      if (
        // check that the api function doesn't take any arguments
        !apiFunction.length &&
        // but the user passed an empty array as the only parameter
        parameters.length === 1 &&
        Array.isArray(parameters[0]) &&
        parameters[0].length === 0
      ) {
        dataReaderObj.current = initializeDataReader(apiFunction as ApiFn<ResponseType>)
      } else {
        // eager initialization for all other cases
        dataReaderObj.current = initializeDataReader(
          apiFunction as ApiFn<ResponseType, ArgTypes>,
          ...parameters
        )
      }
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [apiFunction, ...parameters])

  // state to force re-render
  const [, forceRender] = useState(0)

  const updaterFn = useCallback(
    (...newParameters: ArgTypes) => {
      // update the object ref
      dataReaderObj.current = initializeDataReader(
        apiFunction as ApiFn<ResponseType, ArgTypes>,
        ...newParameters
      )
      // update state to force a re-render
      forceRender((ct) => 1 - ct)
    },
    [apiFunction]
  )

  return [dataReaderObj.current, updaterFn]
}
