import { UseTooltipProps } from "../../../tooltip/src/tooltip.js";
import { TableColumnCtx } from "./defaults.js";
import { TableSortOrder } from "../table/defaults.js";
import "../../../../index.js";
import * as vue from "vue";

//#region ../../packages/components/table/src/table-column/index.d.ts
declare const _default: vue.DefineComponent<{
  type: {
    type: StringConstructor;
    default: string;
  };
  label: StringConstructor;
  className: StringConstructor;
  labelClassName: StringConstructor;
  property: StringConstructor;
  prop: StringConstructor;
  width: {
    type: (NumberConstructor | StringConstructor)[];
    default: string;
  };
  minWidth: {
    type: (NumberConstructor | StringConstructor)[];
    default: string;
  };
  renderHeader: vue.PropType<TableColumnCtx<any>["renderHeader"]>;
  sortable: {
    type: (BooleanConstructor | StringConstructor)[];
    default: boolean;
  };
  sortMethod: vue.PropType<TableColumnCtx<any>["sortMethod"]>;
  sortBy: vue.PropType<TableColumnCtx<any>["sortBy"]>;
  resizable: {
    type: BooleanConstructor;
    default: boolean;
  };
  columnKey: StringConstructor;
  align: StringConstructor;
  headerAlign: StringConstructor;
  showOverflowTooltip: {
    type: vue.PropType<TableColumnCtx<any>["showOverflowTooltip"]>;
    default: undefined;
  };
  tooltipFormatter: vue.PropType<TableColumnCtx<any>["tooltipFormatter"]>;
  fixed: (BooleanConstructor | StringConstructor)[];
  formatter: vue.PropType<TableColumnCtx<any>["formatter"]>;
  selectable: vue.PropType<TableColumnCtx<any>["selectable"]>;
  reserveSelection: BooleanConstructor;
  filterMethod: vue.PropType<TableColumnCtx<any>["filterMethod"]>;
  filteredValue: vue.PropType<TableColumnCtx<any>["filteredValue"]>;
  filters: vue.PropType<TableColumnCtx<any>["filters"]>;
  filterPlacement: StringConstructor;
  filterMultiple: {
    type: BooleanConstructor;
    default: boolean;
  };
  filterClassName: StringConstructor;
  index: vue.PropType<TableColumnCtx<any>["index"]>;
  sortOrders: {
    type: vue.PropType<TableColumnCtx<any>["sortOrders"]>;
    default: () => (string | null)[];
    validator: (val: TableColumnCtx<any>["sortOrders"]) => boolean;
  };
}, void, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<{
  type: {
    type: StringConstructor;
    default: string;
  };
  label: StringConstructor;
  className: StringConstructor;
  labelClassName: StringConstructor;
  property: StringConstructor;
  prop: StringConstructor;
  width: {
    type: (NumberConstructor | StringConstructor)[];
    default: string;
  };
  minWidth: {
    type: (NumberConstructor | StringConstructor)[];
    default: string;
  };
  renderHeader: vue.PropType<TableColumnCtx<any>["renderHeader"]>;
  sortable: {
    type: (BooleanConstructor | StringConstructor)[];
    default: boolean;
  };
  sortMethod: vue.PropType<TableColumnCtx<any>["sortMethod"]>;
  sortBy: vue.PropType<TableColumnCtx<any>["sortBy"]>;
  resizable: {
    type: BooleanConstructor;
    default: boolean;
  };
  columnKey: StringConstructor;
  align: StringConstructor;
  headerAlign: StringConstructor;
  showOverflowTooltip: {
    type: vue.PropType<TableColumnCtx<any>["showOverflowTooltip"]>;
    default: undefined;
  };
  tooltipFormatter: vue.PropType<TableColumnCtx<any>["tooltipFormatter"]>;
  fixed: (BooleanConstructor | StringConstructor)[];
  formatter: vue.PropType<TableColumnCtx<any>["formatter"]>;
  selectable: vue.PropType<TableColumnCtx<any>["selectable"]>;
  reserveSelection: BooleanConstructor;
  filterMethod: vue.PropType<TableColumnCtx<any>["filterMethod"]>;
  filteredValue: vue.PropType<TableColumnCtx<any>["filteredValue"]>;
  filters: vue.PropType<TableColumnCtx<any>["filters"]>;
  filterPlacement: StringConstructor;
  filterMultiple: {
    type: BooleanConstructor;
    default: boolean;
  };
  filterClassName: StringConstructor;
  index: vue.PropType<TableColumnCtx<any>["index"]>;
  sortOrders: {
    type: vue.PropType<TableColumnCtx<any>["sortOrders"]>;
    default: () => (string | null)[];
    validator: (val: TableColumnCtx<any>["sortOrders"]) => boolean;
  };
}>>, {
  type: string;
  minWidth: string | number;
  width: string | number;
  resizable: boolean;
  showOverflowTooltip: boolean | Partial<Pick<UseTooltipProps, "offset" | "appendTo" | "effect" | "placement" | "popperClass" | "enterable" | "popperOptions" | "showArrow" | "transition" | "showAfter" | "hideAfter">> | undefined;
  sortOrders: (TableSortOrder | null)[];
  sortable: string | boolean;
  reserveSelection: boolean;
  filterMultiple: boolean;
}, {}>;
//#endregion
export { _default as default };