import { Translator } from "../../../../hooks/use-locale/index.js";
import { TableColumnCtx } from "../table-column/defaults.js";
import { Store } from "../store/index.js";
import { TableLayout } from "../table-layout.js";
import { DefaultRow, Sort, TableSortOrder } from "../table/defaults.js";
import * as vue from "vue";
import { ComponentInternalInstance, PropType, Ref } from "vue";

//#region ../../packages/components/table/src/table-header/index.d.ts
interface TableHeader extends ComponentInternalInstance {
  state: {
    onColumnsChange: (layout: TableLayout<any>) => void;
    onScrollableChange: (layout: TableLayout<any>) => void;
  };
  filterPanels: Ref<DefaultRow>;
}
interface TableHeaderProps<T extends DefaultRow> {
  fixed: string;
  store: Store<T>;
  border: boolean;
  defaultSort: Sort;
  allowDragLastColumn: boolean;
}
declare const _default: vue.DefineComponent<{
  fixed: {
    type: StringConstructor;
    default: string;
  };
  store: {
    required: true;
    type: PropType<TableHeaderProps<any>["store"]>;
  };
  border: BooleanConstructor;
  defaultSort: {
    type: PropType<TableHeaderProps<any>["defaultSort"]>;
    default: () => {
      prop: string;
      order: string;
    };
  };
  appendFilterPanelTo: {
    type: StringConstructor;
  };
  allowDragLastColumn: {
    type: BooleanConstructor;
  };
}, {
  ns: {
    namespace: vue.ComputedRef<string>;
    b: (blockSuffix?: string) => string;
    e: (element?: string) => string;
    m: (modifier?: string) => string;
    be: (blockSuffix?: string, element?: string) => string;
    em: (element?: string, modifier?: string) => string;
    bm: (blockSuffix?: string, modifier?: string) => string;
    bem: (blockSuffix?: string, element?: string, modifier?: string) => string;
    is: {
      (name: string, state: boolean | undefined): string;
      (name: string): string;
    };
    cssVar: (object: Record<string, string>) => Record<string, string>;
    cssVarName: (name: string) => string;
    cssVarBlock: (object: Record<string, string>) => Record<string, string>;
    cssVarBlockName: (name: string) => string;
  };
  t: Translator;
  filterPanels: Ref<{}>;
  onColumnsChange: (layout: TableLayout<DefaultRow>) => void;
  onScrollableChange: (layout: TableLayout<DefaultRow>) => void;
  columnRows: vue.ComputedRef<TableColumnCtx<any>[][]>;
  getHeaderRowClass: (rowIndex: number) => string;
  getHeaderRowStyle: (rowIndex: number) => any;
  getHeaderCellClass: (rowIndex: number, columnIndex: number, row: any, column: TableColumnCtx<any>) => string;
  getHeaderCellStyle: (rowIndex: number, columnIndex: number, row: any, column: TableColumnCtx<any>) => vue.CSSProperties;
  handleHeaderClick: (event: Event, column: TableColumnCtx<any>) => void;
  handleHeaderContextMenu: (event: Event, column: TableColumnCtx<any>) => void;
  handleMouseDown: (event: MouseEvent, column: TableColumnCtx<any>) => void;
  handleMouseMove: (event: MouseEvent, column: TableColumnCtx<any>) => void;
  handleMouseOut: () => void;
  handleSortClick: (event: Event, column: TableColumnCtx<any>, givenOrder?: TableSortOrder | boolean) => void;
  handleFilterClick: (event: Event) => void;
  isGroup: vue.ComputedRef<boolean>;
  toggleAllSelection: (event: Event) => void;
  saveIndexSelection: vue.Reactive<Map<any, any>>;
  isTableLayoutAuto: boolean;
  theadRef: Ref<any>;
  updateFixedColumnStyle: () => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<{
  fixed: {
    type: StringConstructor;
    default: string;
  };
  store: {
    required: true;
    type: PropType<TableHeaderProps<any>["store"]>;
  };
  border: BooleanConstructor;
  defaultSort: {
    type: PropType<TableHeaderProps<any>["defaultSort"]>;
    default: () => {
      prop: string;
      order: string;
    };
  };
  appendFilterPanelTo: {
    type: StringConstructor;
  };
  allowDragLastColumn: {
    type: BooleanConstructor;
  };
}>>, {
  fixed: string;
  border: boolean;
  defaultSort: Sort;
  allowDragLastColumn: boolean;
}, {}>;
//#endregion
export { TableHeader, TableHeaderProps, _default as default };