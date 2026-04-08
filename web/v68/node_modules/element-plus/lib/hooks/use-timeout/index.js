Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
const require_runtime = require('../../_virtual/_rolldown/runtime.js');
let _vueuse_core = require("@vueuse/core");

//#region ../../packages/hooks/use-timeout/index.ts
function useTimeout() {
	let timeoutHandle;
	const registerTimeout = (fn, delay) => {
		cancelTimeout();
		timeoutHandle = window.setTimeout(fn, delay);
	};
	const cancelTimeout = () => window.clearTimeout(timeoutHandle);
	(0, _vueuse_core.tryOnScopeDispose)(() => cancelTimeout());
	return {
		registerTimeout,
		cancelTimeout
	};
}

//#endregion
exports.useTimeout = useTimeout;
//# sourceMappingURL=index.js.map