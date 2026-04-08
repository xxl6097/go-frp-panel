import * as vue from "vue";
import { ExtractPropTypes, ExtractPublicPropTypes } from "vue";

//#region ../../packages/components/space/src/item.d.ts
declare const spaceItemProps: {
  readonly prefixCls: {
    readonly type: vue.PropType<string>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
};
type SpaceItemProps = ExtractPropTypes<typeof spaceItemProps>;
type SpaceItemPropsPublic = ExtractPublicPropTypes<typeof spaceItemProps>;
declare const SpaceItem: vue.DefineComponent<{
  readonly prefixCls: {
    readonly type: vue.PropType<string>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
}, () => vue.VNode<vue.RendererNode, vue.RendererElement, {
  [key: string]: any;
}>, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<ExtractPropTypes<{
  readonly prefixCls: {
    readonly type: vue.PropType<string>;
    readonly required: false;
    readonly validator: ((val: unknown) => boolean) | undefined;
    __epPropKey: true;
  };
}>>, {}, {}>;
type SpaceItemInstance = InstanceType<typeof SpaceItem> & unknown;
//#endregion
export { SpaceItemInstance, SpaceItemProps, SpaceItemPropsPublic, spaceItemProps };