Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });

//#region ../../packages/utils/i18n.ts
const isKorean = (text) => /([\uAC00-\uD7AF\u3130-\u318F])+/gi.test(text);

//#endregion
exports.isKorean = isKorean;
//# sourceMappingURL=i18n.js.map