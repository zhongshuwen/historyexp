export function ellipsis(input: string, maxLength: number = 20): string {
  if (input.length <= maxLength) {
    return input
  }

  return `${input.substr(0, maxLength)}…`
}
