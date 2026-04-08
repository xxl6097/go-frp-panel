import { computed } from "vue";

//#region ../../packages/components/slider/src/composables/use-marks.ts
const useMarks = (props) => {
	return computed(() => {
		if (!props.marks) return [];
		return Object.keys(props.marks).map(Number.parseFloat).sort((a, b) => a - b).filter((point) => point <= props.max && point >= props.min).map((point) => ({
			point,
			position: (point - props.min) * 100 / (props.max - props.min),
			mark: props.marks[point]
		}));
	});
};

//#endregion
export { useMarks };
//# sourceMappingURL=use-marks.mjs.map