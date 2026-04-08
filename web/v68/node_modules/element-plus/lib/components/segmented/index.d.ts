import { ComponentSize } from "../../constants/size.js";
import { Option } from "./src/types.js";
import { Props, SegmentedEmits, SegmentedInstance, SegmentedProps, SegmentedPropsPublic, defaultProps, segmentedEmits, segmentedProps } from "./src/segmented.js";
import "../../index.js";
import * as vue from "vue";

//#region ../../packages/components/segmented/index.d.ts
declare const ElSegmented: {
  new (...args: any[]): vue.CreateComponentPublicInstance<Readonly<vue.ExtractPropTypes<{
    props: {
      type: vue.PropType<Props>;
      default: () => Required<Props>;
    };
    modelValue: {
      type: vue.PropType<string | number | boolean>;
      default: undefined;
    };
    id: {
      type: vue.PropType<string>;
    };
    size: {
      type: vue.PropType<"" | "default" | "small" | "large">;
    };
    disabled: {
      type: vue.PropType<boolean>;
      default: undefined;
    };
    validateEvent: {
      type: vue.PropType<boolean>;
      default: boolean;
    };
    ariaLabel: {
      type: vue.PropType<string>;
    };
    name: {
      type: vue.PropType<string>;
    };
    block: {
      type: vue.PropType<boolean>;
    };
    direction: {
      type: vue.PropType<"horizontal" | "vertical">;
      default: string;
    };
    options: {
      type: vue.PropType<Option[]>;
      default: () => never[];
    };
  }>> & {
    onChange?: ((val: any) => any) | undefined;
    "onUpdate:modelValue"?: ((val: any) => any) | undefined;
  }, unknown, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
    change: (val: any) => void;
    "update:modelValue": (val: any) => void;
  }, vue.VNodeProps & vue.AllowedComponentProps & vue.ComponentCustomProps & Readonly<vue.ExtractPropTypes<{
    props: {
      type: vue.PropType<Props>;
      default: () => Required<Props>;
    };
    modelValue: {
      type: vue.PropType<string | number | boolean>;
      default: undefined;
    };
    id: {
      type: vue.PropType<string>;
    };
    size: {
      type: vue.PropType<"" | "default" | "small" | "large">;
    };
    disabled: {
      type: vue.PropType<boolean>;
      default: undefined;
    };
    validateEvent: {
      type: vue.PropType<boolean>;
      default: boolean;
    };
    ariaLabel: {
      type: vue.PropType<string>;
    };
    name: {
      type: vue.PropType<string>;
    };
    block: {
      type: vue.PropType<boolean>;
    };
    direction: {
      type: vue.PropType<"horizontal" | "vertical">;
      default: string;
    };
    options: {
      type: vue.PropType<Option[]>;
      default: () => never[];
    };
  }>> & {
    onChange?: ((val: any) => any) | undefined;
    "onUpdate:modelValue"?: ((val: any) => any) | undefined;
  }, {
    props: Props;
    modelValue: string | number | boolean;
    disabled: boolean;
    validateEvent: boolean;
    direction: "vertical" | "horizontal";
    options: Option[];
  }, true, {}, {}, {
    P: {};
    B: {};
    D: {};
    C: {};
    M: {};
    Defaults: {};
  }, Readonly<vue.ExtractPropTypes<{
    props: {
      type: vue.PropType<Props>;
      default: () => Required<Props>;
    };
    modelValue: {
      type: vue.PropType<string | number | boolean>;
      default: undefined;
    };
    id: {
      type: vue.PropType<string>;
    };
    size: {
      type: vue.PropType<"" | "default" | "small" | "large">;
    };
    disabled: {
      type: vue.PropType<boolean>;
      default: undefined;
    };
    validateEvent: {
      type: vue.PropType<boolean>;
      default: boolean;
    };
    ariaLabel: {
      type: vue.PropType<string>;
    };
    name: {
      type: vue.PropType<string>;
    };
    block: {
      type: vue.PropType<boolean>;
    };
    direction: {
      type: vue.PropType<"horizontal" | "vertical">;
      default: string;
    };
    options: {
      type: vue.PropType<Option[]>;
      default: () => never[];
    };
  }>> & {
    onChange?: ((val: any) => any) | undefined;
    "onUpdate:modelValue"?: ((val: any) => any) | undefined;
  }, {}, {}, {}, {}, {
    props: Props;
    modelValue: string | number | boolean;
    disabled: boolean;
    validateEvent: boolean;
    direction: "vertical" | "horizontal";
    options: Option[];
  }>;
  __isFragment?: never;
  __isTeleport?: never;
  __isSuspense?: never;
} & vue.ComponentOptionsBase<Readonly<vue.ExtractPropTypes<{
  props: {
    type: vue.PropType<Props>;
    default: () => Required<Props>;
  };
  modelValue: {
    type: vue.PropType<string | number | boolean>;
    default: undefined;
  };
  id: {
    type: vue.PropType<string>;
  };
  size: {
    type: vue.PropType<"" | "default" | "small" | "large">;
  };
  disabled: {
    type: vue.PropType<boolean>;
    default: undefined;
  };
  validateEvent: {
    type: vue.PropType<boolean>;
    default: boolean;
  };
  ariaLabel: {
    type: vue.PropType<string>;
  };
  name: {
    type: vue.PropType<string>;
  };
  block: {
    type: vue.PropType<boolean>;
  };
  direction: {
    type: vue.PropType<"horizontal" | "vertical">;
    default: string;
  };
  options: {
    type: vue.PropType<Option[]>;
    default: () => never[];
  };
}>> & {
  onChange?: ((val: any) => any) | undefined;
  "onUpdate:modelValue"?: ((val: any) => any) | undefined;
}, unknown, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {
  change: (val: any) => void;
  "update:modelValue": (val: any) => void;
}, string, {
  props: Props;
  modelValue: string | number | boolean;
  disabled: boolean;
  validateEvent: boolean;
  direction: "vertical" | "horizontal";
  options: Option[];
}, {}, string, {}> & vue.VNodeProps & vue.AllowedComponentProps & vue.ComponentCustomProps & (new () => {
  $slots: {
    default?: (props: {
      item: any;
    }) => any;
  };
}) & vue.ObjectPlugin & {
  setPropsDefaults: (defaults: Partial<Partial<{
    props: Props;
    modelValue: string | number | boolean;
    disabled: boolean;
    validateEvent: boolean;
    direction: "vertical" | "horizontal";
    options: Option[];
  }> & Omit<{
    readonly props: Props;
    readonly validateEvent: boolean;
    readonly direction: "vertical" | "horizontal";
    readonly options: Option[];
    readonly modelValue?: (string | number | boolean) | undefined;
    readonly id?: string | undefined;
    readonly size?: ComponentSize | undefined;
    readonly disabled?: boolean | undefined;
    readonly ariaLabel?: string | undefined;
    readonly name?: string | undefined;
    readonly block?: boolean | undefined;
    onChange?: ((val: any) => any) | undefined;
    "onUpdate:modelValue"?: ((val: any) => any) | undefined;
  } & vue.VNodeProps & vue.AllowedComponentProps & vue.ComponentCustomProps & Readonly<vue.ExtractPropTypes<{
    props: {
      type: vue.PropType<Props>;
      default: () => Required<Props>;
    };
    modelValue: {
      type: vue.PropType<string | number | boolean>;
      default: undefined;
    };
    id: {
      type: vue.PropType<string>;
    };
    size: {
      type: vue.PropType<"" | "default" | "small" | "large">;
    };
    disabled: {
      type: vue.PropType<boolean>;
      default: undefined;
    };
    validateEvent: {
      type: vue.PropType<boolean>;
      default: boolean;
    };
    ariaLabel: {
      type: vue.PropType<string>;
    };
    name: {
      type: vue.PropType<string>;
    };
    block: {
      type: vue.PropType<boolean>;
    };
    direction: {
      type: vue.PropType<"horizontal" | "vertical">;
      default: string;
    };
    options: {
      type: vue.PropType<Option[]>;
      default: () => never[];
    };
  }>> & {
    onChange?: ((val: any) => any) | undefined;
    "onUpdate:modelValue"?: ((val: any) => any) | undefined;
  }, "props" | "modelValue" | "disabled" | "validateEvent" | "direction" | "options">>) => void;
} & Record<string, any>;
//#endregion
export { ElSegmented, ElSegmented as default, Props, SegmentedEmits, SegmentedInstance, SegmentedProps, SegmentedPropsPublic, defaultProps, segmentedEmits, segmentedProps };