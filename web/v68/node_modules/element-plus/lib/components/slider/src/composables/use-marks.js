Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
const require_runtime = require('../../../../_virtual/_rolldown/runtime.js');
let vue = require("vue");

//#region ../../packages/components/slider/src/composables/use-marks.ts
const useMarks = (props) => {
	return (0, vue.computed)(() => {
		if (!props.marks) return [];
		return Object.keys(props.marks).map(Number.parseFloat).sort((a, b) => a - b).filter((point) => point <= props.max && point >= props.min).map((point) => ({
			point,
			position: (point - props.min) * 100 / (props.max - props.min),
			mark: props.marks[point]
		}));
	});
};

//#endregion
exports.useMarks = useMarks;
//# sourceMappingURL=use-marks.js.map