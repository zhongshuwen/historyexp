interface SQE {
  query: string
  lowBlockNum?: string
  highBlockNum?: string
}

// @TODO handle invalid queries
export function parseSQE(input: string): SQE {
  // Patch to make transaction/block hash search work
  if (/^[a-zA-Z0-9]+$/.test(input.trim())) {
    return {
      query: input.trim(),
    };
  }

  return {
    query: input.trim(),
  };

  // Disabled for now until we fix it (breaks OR elements)
  // const parts = input.trim().split(' ')
  //   .filter(part => part.indexOf(':') >= 0)
  //   .map(part => part.split(':')); // @TODO what if there multiple ::: in the string?

  // const queryObj: Record<string, string> = Object.fromEntries(parts);
  // const lowBlockNum = queryObj.lowblocknum;
  // const highBlockNum = queryObj.highblocknum;

  // delete queryObj.lowblocknum;
  // delete queryObj.highblocknum;

  // const query = Object.keys(queryObj)
  //   .map(key => `${key}:${queryObj[key]}`)
  //   .join(' ');

  // return {
  //   query,
  //   lowBlockNum,
  //   highBlockNum,
  // };
}
