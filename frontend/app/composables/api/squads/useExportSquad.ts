export function useExportSquad() {
	const downloading = ref(false);

	const downloadBlob = (blob: Blob, filename: string) => {
		const url = URL.createObjectURL(blob);
		const a = document.createElement("a");
		a.href = url;
		a.download = filename;
		document.body.appendChild(a);
		a.click();
		document.body.removeChild(a);
		URL.revokeObjectURL(url);
	};

	const exportAll = async (format: string, myOnly: boolean = true) => {
		downloading.value = true;
		try {
			const res = await $fetch<Blob>(`/api/student-squads/export/${format}?myOnly=${myOnly}`, {
				responseType: "blob",
			});
			const ext = format === "md" ? ".md" : format === "html" ? ".html" : ".pdf";
			downloadBlob(res, `student_squads${ext}`);
		} finally {
			downloading.value = false;
		}
	};

	const exportByID = async (id: string, format: string) => {
		downloading.value = true;
		try {
			const res = await $fetch<Blob>(`/api/student-squads/${id}/export/${format}`, {
				responseType: "blob",
			});
			const ext = format === "md" ? ".md" : format === "html" ? ".html" : ".pdf";
			downloadBlob(res, `squad_${id.slice(0, 8)}${ext}`);
		} finally {
			downloading.value = false;
		}
	};

	return { exportAll, exportByID, downloading };
}
