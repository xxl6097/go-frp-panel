import * as vue from "vue";
import { Ref, VNode } from "vue";

//#region ../../packages/components/slot/src/only-child.d.ts
declare const OnlyChild: vue.DefineComponent<{}, () => VNode<vue.RendererNode, vue.RendererElement, {
  [key: string]: any;
}> | null, {}, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<{}>>, {}, {}>;
type OnlyChildExpose = {
  forwardRef: Ref<HTMLElement>;
};
//#endregion
export { OnlyChild, OnlyChildExpose };