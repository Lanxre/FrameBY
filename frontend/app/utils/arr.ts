export const mergeByProp = (
	a1: any[],
	a2: any[],
	prop: string = "id",
): any[] => {
	const a2Ids = new Set(a2.map((item) => item[prop]));

	const filtered = a1.filter((item) => !a2Ids.has(item[prop]));

	return [...filtered, ...a2];
};

export const removeFromFirstByProp = (
	first: any[],
	second: any[],
	prop: string = "id",
): any[] => {
	const secondValues = new Set(second.map((item) => item[prop]));

	return first.filter((item) => !secondValues.has(item[prop]));
};
