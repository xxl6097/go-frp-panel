import * as vue from "vue";
import { ComponentInternalInstance, VNode } from "vue";

//#region ../../packages/hooks/use-ordered-children/index.d.ts
type ChildEssential = {
  uid: number;
  getVnode: () => VNode;
};
declare const useOrderedChildren: <T extends ChildEssential>(vm: ComponentInternalInstance, childComponentName: string) => {
  children: vue.ShallowRef<T[]>;
  addChild: (child: T) => void;
  removeChild: (child: T) => void;
  ChildrenSorter: vue.DefineComponent<{}, () => VNode<vue.RendererNode, vue.RendererElement, {
    [key: string]: any;
  }> | null, {}, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<{}>>, {}, {}>;
};
//#endregion
export { useOrderedChildren };