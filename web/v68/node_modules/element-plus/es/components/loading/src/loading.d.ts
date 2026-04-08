import { LoadingOptionsResolved, LoadingParentElement } from "./types.js";
import * as vue from "vue";
import { AppContext, VNode } from "vue";

//#region ../../packages/components/loading/src/loading.d.ts
declare function createLoadingComponent(options: LoadingOptionsResolved, appContext: AppContext | null): {
  setText: (text: string | VNode | VNode[]) => void;
  removeElLoadingChild: () => void;
  close: () => void;
  handleAfterLeave: () => void;
  vm: vue.ComponentPublicInstance<{}, {}, {}, {}, {}, {}, {}, {}, false, vue.ComponentOptionsBase<any, any, any, any, any, any, any, any, any, {}, {}, string, {}>, {}, {}>;
  $el: HTMLElement;
  originalPosition: vue.Ref<string>;
  originalOverflow: vue.Ref<string>;
  visible: vue.Ref<boolean>;
  parent: vue.Ref<LoadingParentElement>;
  background: vue.Ref<string>;
  svg: vue.Ref<string>;
  svgViewBox: vue.Ref<string>;
  spinner: vue.Ref<string | boolean>;
  text: vue.Ref<string | VNode<vue.RendererNode, vue.RendererElement, {
    [key: string]: any;
  }> | VNode<vue.RendererNode, vue.RendererElement, {
    [key: string]: any;
  }>[]>;
  fullscreen: vue.Ref<boolean>;
  lock: vue.Ref<boolean>;
  customClass: vue.Ref<string>;
  target: vue.Ref<HTMLElement>;
  beforeClose?: vue.Ref<(() => boolean) | undefined> | undefined;
  closed?: vue.Ref<(() => void) | undefined> | undefined;
};
type LoadingInstance = ReturnType<typeof createLoadingComponent>;
//#endregion
export { LoadingInstance };