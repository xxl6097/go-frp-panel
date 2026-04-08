import { EpPropMergeType } from "../../../utils/vue/props/types.js";
import "../../../utils/index.js";
import { Translator } from "../../../hooks/use-locale/index.js";
import "../../../hooks/index.js";
import { UseTooltipProps } from "../../tooltip/src/tooltip.js";
import { TreeData } from "./store/tree.js";
import { TableColumnCtx } from "./table-column/defaults.js";
import { StoreFilter } from "./store/index.js";
import { TableLayout } from "./table-layout.js";
import { Filter, RenderExpanded, Sort, Table, TableProps, TableSortOrder, TreeProps } from "./table/defaults.js";
import "../../../index.js";
import * as vue from "vue";
import * as lodash from "lodash";

//#region ../../packages/components/table/src/table.vue.d.ts
declare const _default: typeof __VLS_export;
declare const __VLS_export: vue.DefineComponent<{
  data: {
    type: vue.PropType<any[]>;
    default: () => never[];
  };
  size: {
    readonly type: vue.PropType<EpPropMergeType<StringConstructor, "" | "default" | "small" | "large", never>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  width: (NumberConstructor | StringConstructor)[];
  height: (NumberConstructor | StringConstructor)[];
  maxHeight: (NumberConstructor | StringConstructor)[];
  fit: {
    type: BooleanConstructor;
    default: boolean;
  };
  stripe: BooleanConstructor;
  border: BooleanConstructor;
  rowKey: vue.PropType<TableProps<any>["rowKey"]>;
  showHeader: {
    type: BooleanConstructor;
    default: boolean;
  };
  showSummary: BooleanConstructor;
  sumText: StringConstructor;
  summaryMethod: vue.PropType<TableProps<any>["summaryMethod"]>;
  rowClassName: vue.PropType<TableProps<any>["rowClassName"]>;
  rowStyle: vue.PropType<TableProps<any>["rowStyle"]>;
  cellClassName: vue.PropType<TableProps<any>["cellClassName"]>;
  cellStyle: vue.PropType<TableProps<any>["cellStyle"]>;
  headerRowClassName: vue.PropType<TableProps<any>["headerRowClassName"]>;
  headerRowStyle: vue.PropType<TableProps<any>["headerRowStyle"]>;
  headerCellClassName: vue.PropType<TableProps<any>["headerCellClassName"]>;
  headerCellStyle: vue.PropType<TableProps<any>["headerCellStyle"]>;
  highlightCurrentRow: BooleanConstructor;
  currentRowKey: (NumberConstructor | StringConstructor)[];
  emptyText: StringConstructor;
  expandRowKeys: vue.PropType<TableProps<any>["expandRowKeys"]>;
  defaultExpandAll: BooleanConstructor;
  rowExpandable: {
    type: vue.PropType<TableProps<any>["rowExpandable"]>;
  };
  defaultSort: vue.PropType<TableProps<any>["defaultSort"]>;
  tooltipEffect: StringConstructor;
  tooltipOptions: vue.PropType<TableProps<any>["tooltipOptions"]>;
  spanMethod: vue.PropType<TableProps<any>["spanMethod"]>;
  selectOnIndeterminate: {
    type: BooleanConstructor;
    default: boolean;
  };
  indent: {
    type: NumberConstructor;
    default: number;
  };
  treeProps: {
    type: vue.PropType<TableProps<any>["treeProps"]>;
    default: () => {
      hasChildren: string;
      children: string;
      checkStrictly: boolean;
    };
  };
  lazy: BooleanConstructor;
  load: vue.PropType<TableProps<any>["load"]>;
  style: {
    type: vue.PropType<TableProps<any>["style"]>;
    default: () => {};
  };
  className: {
    type: StringConstructor;
    default: string;
  };
  tableLayout: {
    type: vue.PropType<"fixed" | "auto">;
    default: string;
  };
  scrollbarAlwaysOn: BooleanConstructor;
  flexible: BooleanConstructor;
  showOverflowTooltip: {
    type: vue.PropType<TableProps<any>["showOverflowTooltip"]>;
    default: undefined;
  };
  tooltipFormatter: vue.PropType<TableProps<any>["tooltipFormatter"]>;
  appendFilterPanelTo: StringConstructor;
  scrollbarTabindex: {
    type: (NumberConstructor | StringConstructor)[];
    default: undefined;
  };
  allowDragLastColumn: {
    type: BooleanConstructor;
    default: boolean;
  };
  preserveExpandedContent: BooleanConstructor;
  nativeScrollbar: BooleanConstructor;
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
  layout: TableLayout<any>;
  store: {
    mutations: {
      setData(states: {
        _currentRowKey: vue.Ref<string | null>;
        currentRow: vue.Ref<any>;
        expandRowKeys: vue.Ref<string[]>;
        treeData: vue.Ref<Record<string, TreeData>>;
        indent: vue.Ref<number>;
        lazy: vue.Ref<boolean>;
        lazyTreeNodeMap: vue.Ref<Record<string, any[]>>;
        lazyColumnIdentifier: vue.Ref<string>;
        childrenColumnName: vue.Ref<string>;
        checkStrictly: vue.Ref<boolean>;
        expandRows: vue.Ref<any[]>;
        defaultExpandAll: vue.Ref<boolean>;
        tableSize: vue.Ref<any>;
        rowKey: vue.Ref<string | null>;
        data: vue.Ref<any[]>;
        _data: vue.Ref<any[]>;
        isComplex: vue.Ref<boolean>;
        _columns: vue.Ref<TableColumnCtx<any>[]>;
        originColumns: vue.Ref<TableColumnCtx<any>[]>;
        columns: vue.Ref<TableColumnCtx<any>[]>;
        fixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        leafColumns: vue.Ref<TableColumnCtx<any>[]>;
        fixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        updateOrderFns: (() => void)[];
        leafColumnsLength: vue.Ref<number>;
        fixedLeafColumnsLength: vue.Ref<number>;
        rightFixedLeafColumnsLength: vue.Ref<number>;
        isAllSelected: vue.Ref<boolean>;
        selection: vue.Ref<any[]>;
        reserveSelection: vue.Ref<boolean>;
        selectOnIndeterminate: vue.Ref<boolean>;
        selectable: vue.Ref<((row: any, index: number) => boolean) | null>;
        rowExpandable: vue.Ref<((row: any, index: number) => boolean) | null>;
        filters: vue.Ref<StoreFilter>;
        filteredData: vue.Ref<any[] | null>;
        sortingColumn: vue.Ref<TableColumnCtx<any> | null>;
        sortProp: vue.Ref<string | null>;
        sortOrder: vue.Ref<string | number | null>;
        hoverRow: vue.Ref<any>;
      }, data: any[]): void;
      insertColumn(states: {
        _currentRowKey: vue.Ref<string | null>;
        currentRow: vue.Ref<any>;
        expandRowKeys: vue.Ref<string[]>;
        treeData: vue.Ref<Record<string, TreeData>>;
        indent: vue.Ref<number>;
        lazy: vue.Ref<boolean>;
        lazyTreeNodeMap: vue.Ref<Record<string, any[]>>;
        lazyColumnIdentifier: vue.Ref<string>;
        childrenColumnName: vue.Ref<string>;
        checkStrictly: vue.Ref<boolean>;
        expandRows: vue.Ref<any[]>;
        defaultExpandAll: vue.Ref<boolean>;
        tableSize: vue.Ref<any>;
        rowKey: vue.Ref<string | null>;
        data: vue.Ref<any[]>;
        _data: vue.Ref<any[]>;
        isComplex: vue.Ref<boolean>;
        _columns: vue.Ref<TableColumnCtx<any>[]>;
        originColumns: vue.Ref<TableColumnCtx<any>[]>;
        columns: vue.Ref<TableColumnCtx<any>[]>;
        fixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        leafColumns: vue.Ref<TableColumnCtx<any>[]>;
        fixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        updateOrderFns: (() => void)[];
        leafColumnsLength: vue.Ref<number>;
        fixedLeafColumnsLength: vue.Ref<number>;
        rightFixedLeafColumnsLength: vue.Ref<number>;
        isAllSelected: vue.Ref<boolean>;
        selection: vue.Ref<any[]>;
        reserveSelection: vue.Ref<boolean>;
        selectOnIndeterminate: vue.Ref<boolean>;
        selectable: vue.Ref<((row: any, index: number) => boolean) | null>;
        rowExpandable: vue.Ref<((row: any, index: number) => boolean) | null>;
        filters: vue.Ref<StoreFilter>;
        filteredData: vue.Ref<any[] | null>;
        sortingColumn: vue.Ref<TableColumnCtx<any> | null>;
        sortProp: vue.Ref<string | null>;
        sortOrder: vue.Ref<string | number | null>;
        hoverRow: vue.Ref<any>;
      }, column: TableColumnCtx<any>, parent: TableColumnCtx<any>, updateColumnOrder: () => void): void;
      updateColumnOrder(states: {
        _currentRowKey: vue.Ref<string | null>;
        currentRow: vue.Ref<any>;
        expandRowKeys: vue.Ref<string[]>;
        treeData: vue.Ref<Record<string, TreeData>>;
        indent: vue.Ref<number>;
        lazy: vue.Ref<boolean>;
        lazyTreeNodeMap: vue.Ref<Record<string, any[]>>;
        lazyColumnIdentifier: vue.Ref<string>;
        childrenColumnName: vue.Ref<string>;
        checkStrictly: vue.Ref<boolean>;
        expandRows: vue.Ref<any[]>;
        defaultExpandAll: vue.Ref<boolean>;
        tableSize: vue.Ref<any>;
        rowKey: vue.Ref<string | null>;
        data: vue.Ref<any[]>;
        _data: vue.Ref<any[]>;
        isComplex: vue.Ref<boolean>;
        _columns: vue.Ref<TableColumnCtx<any>[]>;
        originColumns: vue.Ref<TableColumnCtx<any>[]>;
        columns: vue.Ref<TableColumnCtx<any>[]>;
        fixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        leafColumns: vue.Ref<TableColumnCtx<any>[]>;
        fixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        updateOrderFns: (() => void)[];
        leafColumnsLength: vue.Ref<number>;
        fixedLeafColumnsLength: vue.Ref<number>;
        rightFixedLeafColumnsLength: vue.Ref<number>;
        isAllSelected: vue.Ref<boolean>;
        selection: vue.Ref<any[]>;
        reserveSelection: vue.Ref<boolean>;
        selectOnIndeterminate: vue.Ref<boolean>;
        selectable: vue.Ref<((row: any, index: number) => boolean) | null>;
        rowExpandable: vue.Ref<((row: any, index: number) => boolean) | null>;
        filters: vue.Ref<StoreFilter>;
        filteredData: vue.Ref<any[] | null>;
        sortingColumn: vue.Ref<TableColumnCtx<any> | null>;
        sortProp: vue.Ref<string | null>;
        sortOrder: vue.Ref<string | number | null>;
        hoverRow: vue.Ref<any>;
      }, column: TableColumnCtx<any>): void;
      removeColumn(states: {
        _currentRowKey: vue.Ref<string | null>;
        currentRow: vue.Ref<any>;
        expandRowKeys: vue.Ref<string[]>;
        treeData: vue.Ref<Record<string, TreeData>>;
        indent: vue.Ref<number>;
        lazy: vue.Ref<boolean>;
        lazyTreeNodeMap: vue.Ref<Record<string, any[]>>;
        lazyColumnIdentifier: vue.Ref<string>;
        childrenColumnName: vue.Ref<string>;
        checkStrictly: vue.Ref<boolean>;
        expandRows: vue.Ref<any[]>;
        defaultExpandAll: vue.Ref<boolean>;
        tableSize: vue.Ref<any>;
        rowKey: vue.Ref<string | null>;
        data: vue.Ref<any[]>;
        _data: vue.Ref<any[]>;
        isComplex: vue.Ref<boolean>;
        _columns: vue.Ref<TableColumnCtx<any>[]>;
        originColumns: vue.Ref<TableColumnCtx<any>[]>;
        columns: vue.Ref<TableColumnCtx<any>[]>;
        fixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        leafColumns: vue.Ref<TableColumnCtx<any>[]>;
        fixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        updateOrderFns: (() => void)[];
        leafColumnsLength: vue.Ref<number>;
        fixedLeafColumnsLength: vue.Ref<number>;
        rightFixedLeafColumnsLength: vue.Ref<number>;
        isAllSelected: vue.Ref<boolean>;
        selection: vue.Ref<any[]>;
        reserveSelection: vue.Ref<boolean>;
        selectOnIndeterminate: vue.Ref<boolean>;
        selectable: vue.Ref<((row: any, index: number) => boolean) | null>;
        rowExpandable: vue.Ref<((row: any, index: number) => boolean) | null>;
        filters: vue.Ref<StoreFilter>;
        filteredData: vue.Ref<any[] | null>;
        sortingColumn: vue.Ref<TableColumnCtx<any> | null>;
        sortProp: vue.Ref<string | null>;
        sortOrder: vue.Ref<string | number | null>;
        hoverRow: vue.Ref<any>;
      }, column: TableColumnCtx<any>, parent: TableColumnCtx<any>, updateColumnOrder: () => void): void;
      sort(states: {
        _currentRowKey: vue.Ref<string | null>;
        currentRow: vue.Ref<any>;
        expandRowKeys: vue.Ref<string[]>;
        treeData: vue.Ref<Record<string, TreeData>>;
        indent: vue.Ref<number>;
        lazy: vue.Ref<boolean>;
        lazyTreeNodeMap: vue.Ref<Record<string, any[]>>;
        lazyColumnIdentifier: vue.Ref<string>;
        childrenColumnName: vue.Ref<string>;
        checkStrictly: vue.Ref<boolean>;
        expandRows: vue.Ref<any[]>;
        defaultExpandAll: vue.Ref<boolean>;
        tableSize: vue.Ref<any>;
        rowKey: vue.Ref<string | null>;
        data: vue.Ref<any[]>;
        _data: vue.Ref<any[]>;
        isComplex: vue.Ref<boolean>;
        _columns: vue.Ref<TableColumnCtx<any>[]>;
        originColumns: vue.Ref<TableColumnCtx<any>[]>;
        columns: vue.Ref<TableColumnCtx<any>[]>;
        fixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        leafColumns: vue.Ref<TableColumnCtx<any>[]>;
        fixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        updateOrderFns: (() => void)[];
        leafColumnsLength: vue.Ref<number>;
        fixedLeafColumnsLength: vue.Ref<number>;
        rightFixedLeafColumnsLength: vue.Ref<number>;
        isAllSelected: vue.Ref<boolean>;
        selection: vue.Ref<any[]>;
        reserveSelection: vue.Ref<boolean>;
        selectOnIndeterminate: vue.Ref<boolean>;
        selectable: vue.Ref<((row: any, index: number) => boolean) | null>;
        rowExpandable: vue.Ref<((row: any, index: number) => boolean) | null>;
        filters: vue.Ref<StoreFilter>;
        filteredData: vue.Ref<any[] | null>;
        sortingColumn: vue.Ref<TableColumnCtx<any> | null>;
        sortProp: vue.Ref<string | null>;
        sortOrder: vue.Ref<string | number | null>;
        hoverRow: vue.Ref<any>;
      }, options: Sort): void;
      changeSortCondition(states: {
        _currentRowKey: vue.Ref<string | null>;
        currentRow: vue.Ref<any>;
        expandRowKeys: vue.Ref<string[]>;
        treeData: vue.Ref<Record<string, TreeData>>;
        indent: vue.Ref<number>;
        lazy: vue.Ref<boolean>;
        lazyTreeNodeMap: vue.Ref<Record<string, any[]>>;
        lazyColumnIdentifier: vue.Ref<string>;
        childrenColumnName: vue.Ref<string>;
        checkStrictly: vue.Ref<boolean>;
        expandRows: vue.Ref<any[]>;
        defaultExpandAll: vue.Ref<boolean>;
        tableSize: vue.Ref<any>;
        rowKey: vue.Ref<string | null>;
        data: vue.Ref<any[]>;
        _data: vue.Ref<any[]>;
        isComplex: vue.Ref<boolean>;
        _columns: vue.Ref<TableColumnCtx<any>[]>;
        originColumns: vue.Ref<TableColumnCtx<any>[]>;
        columns: vue.Ref<TableColumnCtx<any>[]>;
        fixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        leafColumns: vue.Ref<TableColumnCtx<any>[]>;
        fixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        updateOrderFns: (() => void)[];
        leafColumnsLength: vue.Ref<number>;
        fixedLeafColumnsLength: vue.Ref<number>;
        rightFixedLeafColumnsLength: vue.Ref<number>;
        isAllSelected: vue.Ref<boolean>;
        selection: vue.Ref<any[]>;
        reserveSelection: vue.Ref<boolean>;
        selectOnIndeterminate: vue.Ref<boolean>;
        selectable: vue.Ref<((row: any, index: number) => boolean) | null>;
        rowExpandable: vue.Ref<((row: any, index: number) => boolean) | null>;
        filters: vue.Ref<StoreFilter>;
        filteredData: vue.Ref<any[] | null>;
        sortingColumn: vue.Ref<TableColumnCtx<any> | null>;
        sortProp: vue.Ref<string | null>;
        sortOrder: vue.Ref<string | number | null>;
        hoverRow: vue.Ref<any>;
      }, options: Sort): void;
      filterChange(_states: {
        _currentRowKey: vue.Ref<string | null>;
        currentRow: vue.Ref<any>;
        expandRowKeys: vue.Ref<string[]>;
        treeData: vue.Ref<Record<string, TreeData>>;
        indent: vue.Ref<number>;
        lazy: vue.Ref<boolean>;
        lazyTreeNodeMap: vue.Ref<Record<string, any[]>>;
        lazyColumnIdentifier: vue.Ref<string>;
        childrenColumnName: vue.Ref<string>;
        checkStrictly: vue.Ref<boolean>;
        expandRows: vue.Ref<any[]>;
        defaultExpandAll: vue.Ref<boolean>;
        tableSize: vue.Ref<any>;
        rowKey: vue.Ref<string | null>;
        data: vue.Ref<any[]>;
        _data: vue.Ref<any[]>;
        isComplex: vue.Ref<boolean>;
        _columns: vue.Ref<TableColumnCtx<any>[]>;
        originColumns: vue.Ref<TableColumnCtx<any>[]>;
        columns: vue.Ref<TableColumnCtx<any>[]>;
        fixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        leafColumns: vue.Ref<TableColumnCtx<any>[]>;
        fixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        updateOrderFns: (() => void)[];
        leafColumnsLength: vue.Ref<number>;
        fixedLeafColumnsLength: vue.Ref<number>;
        rightFixedLeafColumnsLength: vue.Ref<number>;
        isAllSelected: vue.Ref<boolean>;
        selection: vue.Ref<any[]>;
        reserveSelection: vue.Ref<boolean>;
        selectOnIndeterminate: vue.Ref<boolean>;
        selectable: vue.Ref<((row: any, index: number) => boolean) | null>;
        rowExpandable: vue.Ref<((row: any, index: number) => boolean) | null>;
        filters: vue.Ref<StoreFilter>;
        filteredData: vue.Ref<any[] | null>;
        sortingColumn: vue.Ref<TableColumnCtx<any> | null>;
        sortProp: vue.Ref<string | null>;
        sortOrder: vue.Ref<string | number | null>;
        hoverRow: vue.Ref<any>;
      }, options: Filter<any>): void;
      toggleAllSelection(): void;
      rowSelectedChanged(_states: {
        _currentRowKey: vue.Ref<string | null>;
        currentRow: vue.Ref<any>;
        expandRowKeys: vue.Ref<string[]>;
        treeData: vue.Ref<Record<string, TreeData>>;
        indent: vue.Ref<number>;
        lazy: vue.Ref<boolean>;
        lazyTreeNodeMap: vue.Ref<Record<string, any[]>>;
        lazyColumnIdentifier: vue.Ref<string>;
        childrenColumnName: vue.Ref<string>;
        checkStrictly: vue.Ref<boolean>;
        expandRows: vue.Ref<any[]>;
        defaultExpandAll: vue.Ref<boolean>;
        tableSize: vue.Ref<any>;
        rowKey: vue.Ref<string | null>;
        data: vue.Ref<any[]>;
        _data: vue.Ref<any[]>;
        isComplex: vue.Ref<boolean>;
        _columns: vue.Ref<TableColumnCtx<any>[]>;
        originColumns: vue.Ref<TableColumnCtx<any>[]>;
        columns: vue.Ref<TableColumnCtx<any>[]>;
        fixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        leafColumns: vue.Ref<TableColumnCtx<any>[]>;
        fixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        updateOrderFns: (() => void)[];
        leafColumnsLength: vue.Ref<number>;
        fixedLeafColumnsLength: vue.Ref<number>;
        rightFixedLeafColumnsLength: vue.Ref<number>;
        isAllSelected: vue.Ref<boolean>;
        selection: vue.Ref<any[]>;
        reserveSelection: vue.Ref<boolean>;
        selectOnIndeterminate: vue.Ref<boolean>;
        selectable: vue.Ref<((row: any, index: number) => boolean) | null>;
        rowExpandable: vue.Ref<((row: any, index: number) => boolean) | null>;
        filters: vue.Ref<StoreFilter>;
        filteredData: vue.Ref<any[] | null>;
        sortingColumn: vue.Ref<TableColumnCtx<any> | null>;
        sortProp: vue.Ref<string | null>;
        sortOrder: vue.Ref<string | number | null>;
        hoverRow: vue.Ref<any>;
      }, row: any): void;
      setHoverRow(states: {
        _currentRowKey: vue.Ref<string | null>;
        currentRow: vue.Ref<any>;
        expandRowKeys: vue.Ref<string[]>;
        treeData: vue.Ref<Record<string, TreeData>>;
        indent: vue.Ref<number>;
        lazy: vue.Ref<boolean>;
        lazyTreeNodeMap: vue.Ref<Record<string, any[]>>;
        lazyColumnIdentifier: vue.Ref<string>;
        childrenColumnName: vue.Ref<string>;
        checkStrictly: vue.Ref<boolean>;
        expandRows: vue.Ref<any[]>;
        defaultExpandAll: vue.Ref<boolean>;
        tableSize: vue.Ref<any>;
        rowKey: vue.Ref<string | null>;
        data: vue.Ref<any[]>;
        _data: vue.Ref<any[]>;
        isComplex: vue.Ref<boolean>;
        _columns: vue.Ref<TableColumnCtx<any>[]>;
        originColumns: vue.Ref<TableColumnCtx<any>[]>;
        columns: vue.Ref<TableColumnCtx<any>[]>;
        fixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        leafColumns: vue.Ref<TableColumnCtx<any>[]>;
        fixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        updateOrderFns: (() => void)[];
        leafColumnsLength: vue.Ref<number>;
        fixedLeafColumnsLength: vue.Ref<number>;
        rightFixedLeafColumnsLength: vue.Ref<number>;
        isAllSelected: vue.Ref<boolean>;
        selection: vue.Ref<any[]>;
        reserveSelection: vue.Ref<boolean>;
        selectOnIndeterminate: vue.Ref<boolean>;
        selectable: vue.Ref<((row: any, index: number) => boolean) | null>;
        rowExpandable: vue.Ref<((row: any, index: number) => boolean) | null>;
        filters: vue.Ref<StoreFilter>;
        filteredData: vue.Ref<any[] | null>;
        sortingColumn: vue.Ref<TableColumnCtx<any> | null>;
        sortProp: vue.Ref<string | null>;
        sortOrder: vue.Ref<string | number | null>;
        hoverRow: vue.Ref<any>;
      }, row: any): void;
      setCurrentRow(_states: {
        _currentRowKey: vue.Ref<string | null>;
        currentRow: vue.Ref<any>;
        expandRowKeys: vue.Ref<string[]>;
        treeData: vue.Ref<Record<string, TreeData>>;
        indent: vue.Ref<number>;
        lazy: vue.Ref<boolean>;
        lazyTreeNodeMap: vue.Ref<Record<string, any[]>>;
        lazyColumnIdentifier: vue.Ref<string>;
        childrenColumnName: vue.Ref<string>;
        checkStrictly: vue.Ref<boolean>;
        expandRows: vue.Ref<any[]>;
        defaultExpandAll: vue.Ref<boolean>;
        tableSize: vue.Ref<any>;
        rowKey: vue.Ref<string | null>;
        data: vue.Ref<any[]>;
        _data: vue.Ref<any[]>;
        isComplex: vue.Ref<boolean>;
        _columns: vue.Ref<TableColumnCtx<any>[]>;
        originColumns: vue.Ref<TableColumnCtx<any>[]>;
        columns: vue.Ref<TableColumnCtx<any>[]>;
        fixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedColumns: vue.Ref<TableColumnCtx<any>[]>;
        leafColumns: vue.Ref<TableColumnCtx<any>[]>;
        fixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        rightFixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
        updateOrderFns: (() => void)[];
        leafColumnsLength: vue.Ref<number>;
        fixedLeafColumnsLength: vue.Ref<number>;
        rightFixedLeafColumnsLength: vue.Ref<number>;
        isAllSelected: vue.Ref<boolean>;
        selection: vue.Ref<any[]>;
        reserveSelection: vue.Ref<boolean>;
        selectOnIndeterminate: vue.Ref<boolean>;
        selectable: vue.Ref<((row: any, index: number) => boolean) | null>;
        rowExpandable: vue.Ref<((row: any, index: number) => boolean) | null>;
        filters: vue.Ref<StoreFilter>;
        filteredData: vue.Ref<any[] | null>;
        sortingColumn: vue.Ref<TableColumnCtx<any> | null>;
        sortProp: vue.Ref<string | null>;
        sortOrder: vue.Ref<string | number | null>;
        hoverRow: vue.Ref<any>;
      }, row: any): void;
    };
    commit: (name: "sort" | "setData" | "insertColumn" | "updateColumnOrder" | "removeColumn" | "changeSortCondition" | "filterChange" | "toggleAllSelection" | "rowSelectedChanged" | "setHoverRow" | "setCurrentRow", ...args: any[]) => void;
    updateTableScrollY: () => void;
    assertRowKey: () => void;
    updateColumns: () => void;
    scheduleLayout: (needUpdateColumns?: boolean, immediate?: boolean) => void;
    isSelected: (row: any) => boolean;
    clearSelection: () => void;
    cleanSelection: () => void;
    getSelectionRows: () => any[];
    toggleRowSelection: (row: any, selected?: boolean, emitChange?: boolean, ignoreSelectable?: boolean) => void;
    _toggleAllSelection: () => void;
    toggleAllSelection: (() => void) | null;
    updateAllSelected: () => void;
    updateFilters: (column: TableColumnCtx<any>, values: string[]) => Record<string, string[]>;
    updateCurrentRow: (_currentRow: any) => void;
    updateSort: (column: TableColumnCtx<any> | null, prop: string | null, order: TableSortOrder | null) => void;
    execFilter: () => void;
    execSort: () => void;
    execQuery: (ignore?: {
      filter: boolean;
    } | undefined) => void;
    clearFilter: (columnKeys?: string[] | string) => void;
    clearSort: () => void;
    toggleRowExpansion: (row: any, expanded?: boolean) => void;
    setExpandRowKeysAdapter: (val: string[]) => void;
    setCurrentRowKey: (key: string) => void;
    toggleRowExpansionAdapter: (row: any, expanded?: boolean) => void;
    isRowExpanded: (row: any) => boolean;
    updateExpandRows: () => void;
    updateCurrentRowData: () => void;
    loadOrToggle: (row: any) => void;
    updateTreeData: (ifChangeExpandRowKeys?: boolean, ifExpandAll?: boolean) => void;
    updateKeyChildren: (key: string, data: any[]) => void;
    states: {
      _currentRowKey: vue.Ref<string | null>;
      currentRow: vue.Ref<any>;
      expandRowKeys: vue.Ref<string[]>;
      treeData: vue.Ref<Record<string, TreeData>>;
      indent: vue.Ref<number>;
      lazy: vue.Ref<boolean>;
      lazyTreeNodeMap: vue.Ref<Record<string, any[]>>;
      lazyColumnIdentifier: vue.Ref<string>;
      childrenColumnName: vue.Ref<string>;
      checkStrictly: vue.Ref<boolean>;
      expandRows: vue.Ref<any[]>;
      defaultExpandAll: vue.Ref<boolean>;
      tableSize: vue.Ref<any>;
      rowKey: vue.Ref<string | null>;
      data: vue.Ref<any[]>;
      _data: vue.Ref<any[]>;
      isComplex: vue.Ref<boolean>;
      _columns: vue.Ref<TableColumnCtx<any>[]>;
      originColumns: vue.Ref<TableColumnCtx<any>[]>;
      columns: vue.Ref<TableColumnCtx<any>[]>;
      fixedColumns: vue.Ref<TableColumnCtx<any>[]>;
      rightFixedColumns: vue.Ref<TableColumnCtx<any>[]>;
      leafColumns: vue.Ref<TableColumnCtx<any>[]>;
      fixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
      rightFixedLeafColumns: vue.Ref<TableColumnCtx<any>[]>;
      updateOrderFns: (() => void)[];
      leafColumnsLength: vue.Ref<number>;
      fixedLeafColumnsLength: vue.Ref<number>;
      rightFixedLeafColumnsLength: vue.Ref<number>;
      isAllSelected: vue.Ref<boolean>;
      selection: vue.Ref<any[]>;
      reserveSelection: vue.Ref<boolean>;
      selectOnIndeterminate: vue.Ref<boolean>;
      selectable: vue.Ref<((row: any, index: number) => boolean) | null>;
      rowExpandable: vue.Ref<((row: any, index: number) => boolean) | null>;
      filters: vue.Ref<StoreFilter>;
      filteredData: vue.Ref<any[] | null>;
      sortingColumn: vue.Ref<TableColumnCtx<any> | null>;
      sortProp: vue.Ref<string | null>;
      sortOrder: vue.Ref<string | number | null>;
      hoverRow: vue.Ref<any>;
    };
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
  };
  columns: vue.ComputedRef<TableColumnCtx<any>[]>;
  handleHeaderFooterMousewheel: (_event: WheelEvent, data: any) => void;
  handleMouseLeave: () => void;
  tableId: string;
  tableSize: vue.ComputedRef<"" | "default" | "small" | "large">;
  isHidden: vue.Ref<boolean>;
  isEmpty: vue.ComputedRef<boolean>;
  renderExpanded: vue.Ref<RenderExpanded<any> | null>;
  resizeProxyVisible: vue.Ref<boolean>;
  resizeState: vue.Ref<{
    width: null | number;
    height: null | number;
    headerHeight: null | number;
  }>;
  isGroup: vue.Ref<boolean>;
  bodyWidth: vue.ComputedRef<string>;
  tableBodyStyles: vue.ComputedRef<{
    width: string;
  }>;
  emptyBlockStyle: vue.ComputedRef<{
    width: string;
    height: string;
  } | undefined>;
  debouncedUpdateLayout: lodash.DebouncedFunc<() => void>;
  /**
   * @description used in single selection Table, set a certain row selected. If called without any parameter, it will clear selection
   */
  setCurrentRow: (row: any) => void;
  /**
   * @description returns the currently selected rows
   */
  getSelectionRows: () => any[];
  /**
   * @description used in multiple selection Table, toggle if a certain row is selected. With the second parameter, you can directly set if this row is selected
   */
  toggleRowSelection: (row: any, selected?: boolean, ignoreSelectable?: boolean) => void;
  /**
   * @description used in multiple selection Table, clear user selection
   */
  clearSelection: () => void;
  /**
   * @description clear filters of the columns whose `columnKey` are passed in. If no params, clear all filters
   */
  clearFilter: (columnKeys?: string[] | string) => void;
  /**
   * @description used in multiple selection Table, toggle select all and deselect all
   */
  toggleAllSelection: () => void;
  /**
   * @description used in expandable Table or tree Table, toggle if a certain row is expanded. With the second parameter, you can directly set if this row is expanded or collapsed
   */
  toggleRowExpansion: (row: any, expanded?: boolean) => void;
  /**
   * @description clear sorting, restore data to the original order
   */
  clearSort: () => void;
  /**
   * @description refresh the layout of Table. When the visibility of Table changes, you may need to call this method to get a correct layout
   */
  doLayout: () => void;
  /**
   * @description sort Table manually. Property `prop` is used to set sort column, property `order` is used to set sort order
   */
  sort: (prop: string, order: string) => void;
  /**
   * @description used in lazy Table, must set `rowKey`, update key children
   */
  updateKeyChildren: (key: string, data: any[]) => void;
  t: Translator;
  setDragVisible: (visible: boolean) => void;
  context: Table<any>;
  computedSumText: vue.ComputedRef<string>;
  computedEmptyText: vue.ComputedRef<string>;
  computedTooltipEffect: vue.ComputedRef<string | undefined>;
  computedTooltipOptions: vue.ComputedRef<Partial<Pick<UseTooltipProps, "offset" | "appendTo" | "effect" | "placement" | "popperClass" | "enterable" | "popperOptions" | "showArrow" | "transition" | "showAfter" | "hideAfter">> | undefined>;
  tableLayout: vue.ComputedRef<"fixed" | "auto">;
  scrollbarViewStyle: {
    display: string;
    verticalAlign: string;
  };
  scrollbarStyle: vue.ComputedRef<{
    height: string;
    maxHeight?: undefined;
  } | {
    maxHeight: string;
    height?: undefined;
  } | {
    height?: undefined;
    maxHeight?: undefined;
  }>;
  scrollBarRef: vue.Ref<any>;
  /**
   * @description scrolls to a particular set of coordinates
   */
  scrollTo: (options: ScrollToOptions | number, yCoord?: number) => void;
  /**
   * @description set horizontal scroll position
   */
  setScrollLeft: (left?: number) => void;
  /**
   * @description set vertical scroll position
   */
  setScrollTop: (top?: number) => void;
  /**
   * @description whether to allow drag the last column
   */
  allowDragLastColumn: boolean;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, ("scroll" | "select" | "select-all" | "expand-change" | "current-change" | "selection-change" | "cell-mouse-enter" | "cell-mouse-leave" | "cell-contextmenu" | "cell-click" | "cell-dblclick" | "row-click" | "row-contextmenu" | "row-dblclick" | "header-click" | "header-contextmenu" | "sort-change" | "filter-change" | "header-dragend")[], "scroll" | "select" | "select-all" | "expand-change" | "current-change" | "selection-change" | "cell-mouse-enter" | "cell-mouse-leave" | "cell-contextmenu" | "cell-click" | "cell-dblclick" | "row-click" | "row-contextmenu" | "row-dblclick" | "header-click" | "header-contextmenu" | "sort-change" | "filter-change" | "header-dragend", vue.PublicProps, Readonly<vue.ExtractPropTypes<{
  data: {
    type: vue.PropType<any[]>;
    default: () => never[];
  };
  size: {
    readonly type: vue.PropType<EpPropMergeType<StringConstructor, "" | "default" | "small" | "large", never>>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
  width: (NumberConstructor | StringConstructor)[];
  height: (NumberConstructor | StringConstructor)[];
  maxHeight: (NumberConstructor | StringConstructor)[];
  fit: {
    type: BooleanConstructor;
    default: boolean;
  };
  stripe: BooleanConstructor;
  border: BooleanConstructor;
  rowKey: vue.PropType<TableProps<any>["rowKey"]>;
  showHeader: {
    type: BooleanConstructor;
    default: boolean;
  };
  showSummary: BooleanConstructor;
  sumText: StringConstructor;
  summaryMethod: vue.PropType<TableProps<any>["summaryMethod"]>;
  rowClassName: vue.PropType<TableProps<any>["rowClassName"]>;
  rowStyle: vue.PropType<TableProps<any>["rowStyle"]>;
  cellClassName: vue.PropType<TableProps<any>["cellClassName"]>;
  cellStyle: vue.PropType<TableProps<any>["cellStyle"]>;
  headerRowClassName: vue.PropType<TableProps<any>["headerRowClassName"]>;
  headerRowStyle: vue.PropType<TableProps<any>["headerRowStyle"]>;
  headerCellClassName: vue.PropType<TableProps<any>["headerCellClassName"]>;
  headerCellStyle: vue.PropType<TableProps<any>["headerCellStyle"]>;
  highlightCurrentRow: BooleanConstructor;
  currentRowKey: (NumberConstructor | StringConstructor)[];
  emptyText: StringConstructor;
  expandRowKeys: vue.PropType<TableProps<any>["expandRowKeys"]>;
  defaultExpandAll: BooleanConstructor;
  rowExpandable: {
    type: vue.PropType<TableProps<any>["rowExpandable"]>;
  };
  defaultSort: vue.PropType<TableProps<any>["defaultSort"]>;
  tooltipEffect: StringConstructor;
  tooltipOptions: vue.PropType<TableProps<any>["tooltipOptions"]>;
  spanMethod: vue.PropType<TableProps<any>["spanMethod"]>;
  selectOnIndeterminate: {
    type: BooleanConstructor;
    default: boolean;
  };
  indent: {
    type: NumberConstructor;
    default: number;
  };
  treeProps: {
    type: vue.PropType<TableProps<any>["treeProps"]>;
    default: () => {
      hasChildren: string;
      children: string;
      checkStrictly: boolean;
    };
  };
  lazy: BooleanConstructor;
  load: vue.PropType<TableProps<any>["load"]>;
  style: {
    type: vue.PropType<TableProps<any>["style"]>;
    default: () => {};
  };
  className: {
    type: StringConstructor;
    default: string;
  };
  tableLayout: {
    type: vue.PropType<"fixed" | "auto">;
    default: string;
  };
  scrollbarAlwaysOn: BooleanConstructor;
  flexible: BooleanConstructor;
  showOverflowTooltip: {
    type: vue.PropType<TableProps<any>["showOverflowTooltip"]>;
    default: undefined;
  };
  tooltipFormatter: vue.PropType<TableProps<any>["tooltipFormatter"]>;
  appendFilterPanelTo: StringConstructor;
  scrollbarTabindex: {
    type: (NumberConstructor | StringConstructor)[];
    default: undefined;
  };
  allowDragLastColumn: {
    type: BooleanConstructor;
    default: boolean;
  };
  preserveExpandedContent: BooleanConstructor;
  nativeScrollbar: BooleanConstructor;
}>> & {
  onScroll?: ((...args: any[]) => any) | undefined;
  onSelect?: ((...args: any[]) => any) | undefined;
  "onExpand-change"?: ((...args: any[]) => any) | undefined;
  "onCurrent-change"?: ((...args: any[]) => any) | undefined;
  "onSelect-all"?: ((...args: any[]) => any) | undefined;
  "onSelection-change"?: ((...args: any[]) => any) | undefined;
  "onCell-mouse-enter"?: ((...args: any[]) => any) | undefined;
  "onCell-mouse-leave"?: ((...args: any[]) => any) | undefined;
  "onCell-contextmenu"?: ((...args: any[]) => any) | undefined;
  "onCell-click"?: ((...args: any[]) => any) | undefined;
  "onCell-dblclick"?: ((...args: any[]) => any) | undefined;
  "onRow-click"?: ((...args: any[]) => any) | undefined;
  "onRow-contextmenu"?: ((...args: any[]) => any) | undefined;
  "onRow-dblclick"?: ((...args: any[]) => any) | undefined;
  "onHeader-click"?: ((...args: any[]) => any) | undefined;
  "onHeader-contextmenu"?: ((...args: any[]) => any) | undefined;
  "onSort-change"?: ((...args: any[]) => any) | undefined;
  "onFilter-change"?: ((...args: any[]) => any) | undefined;
  "onHeader-dragend"?: ((...args: any[]) => any) | undefined;
}, {
  style: vue.StyleValue;
  border: boolean;
  className: string;
  tableLayout: "fixed" | "auto";
  data: any[];
  fit: boolean;
  lazy: boolean;
  scrollbarAlwaysOn: boolean;
  allowDragLastColumn: boolean;
  stripe: boolean;
  treeProps: TreeProps | undefined;
  showOverflowTooltip: boolean | Partial<Pick<UseTooltipProps, "offset" | "appendTo" | "effect" | "placement" | "popperClass" | "enterable" | "popperOptions" | "showArrow" | "transition" | "showAfter" | "hideAfter">> | undefined;
  showHeader: boolean;
  showSummary: boolean;
  highlightCurrentRow: boolean;
  defaultExpandAll: boolean;
  selectOnIndeterminate: boolean;
  indent: number;
  flexible: boolean;
  scrollbarTabindex: string | number;
  preserveExpandedContent: boolean;
  nativeScrollbar: boolean;
}, {}>;
//#endregion
export { _default };