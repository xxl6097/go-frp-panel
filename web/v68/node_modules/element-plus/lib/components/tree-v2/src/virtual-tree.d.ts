import "../../../utils/index.js";
import "../../checkbox/index.js";
import "../../tree/src/tree.type.js";
import "./types.js";
import { InjectionKey } from "vue";

//#region ../../packages/components/tree-v2/src/virtual-tree.d.ts
declare enum TreeOptionsEnum {
  KEY = "id",
  LABEL = "label",
  CHILDREN = "children",
  DISABLED = "disabled",
  CLASS = ""
}
//#endregion
export { TreeOptionsEnum };