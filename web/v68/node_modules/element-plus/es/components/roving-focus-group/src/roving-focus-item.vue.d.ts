import * as vue from "vue";

//#region ../../packages/components/roving-focus-group/src/roving-focus-item.vue.d.ts
declare const _default: typeof __VLS_export;
declare const __VLS_export: vue.DefineComponent<{
  focusable: {
    type: BooleanConstructor;
    default: boolean;
  };
  active: BooleanConstructor;
}, {
  id: vue.Ref<string>;
  handleKeydown: (event: Event) => void;
  handleFocus: (event: Event) => void;
  handleMousedown: (event: Event) => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, ("focus" | "keydown" | "mousedown")[], "focus" | "keydown" | "mousedown", vue.PublicProps, Readonly<vue.ExtractPropTypes<{
  focusable: {
    type: BooleanConstructor;
    default: boolean;
  };
  active: BooleanConstructor;
}>> & {
  onFocus?: ((...args: any[]) => any) | undefined;
  onKeydown?: ((...args: any[]) => any) | undefined;
  onMousedown?: ((...args: any[]) => any) | undefined;
}, {
  active: boolean;
  focusable: boolean;
}, {}>;
//#endregion
export { _default };