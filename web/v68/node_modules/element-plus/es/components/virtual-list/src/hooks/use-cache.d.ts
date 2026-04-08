import * as vue from "vue";
import * as lodash from "lodash";
import * as memoize_one0 from "memoize-one";

//#region ../../packages/components/virtual-list/src/hooks/use-cache.d.ts
declare const useCache: <T>() => vue.ComputedRef<(((_: any, __: any, ___: any) => Record<string, T>) & lodash.MemoizedFunction) | memoize_one0.MemoizedFn<(_: any, __: any, ___: any) => Record<string, T>>>;
//#endregion
export { useCache };