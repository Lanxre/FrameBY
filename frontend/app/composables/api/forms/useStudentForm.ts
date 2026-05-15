import { reactive, ref } from "vue";
import { $api } from "../useApi";

export interface StudentFormData {
	fullName: string;
	university: { id: number | string; name: string } | null;
	specialty: string;
	grade: string;
	phone: string;
}

export function useStudentForm() {
	const isLoading = ref(false);
	const showConfirm = ref(false);
	const errorMessage = ref("");

	const { notify } = useNotificationStore();

	const form = reactive<StudentFormData>({
		fullName: "",
		university: null,
		specialty: "",
		grade: "",
		phone: "",
	});

	const validate = () => {
		if (
			!form.fullName ||
			!form.university ||
			!form.specialty ||
			!form.grade ||
			!form.phone
		) {
			errorMessage.value = "Заполните все поля";
			return false;
		}

		const gradeNum = Number(form.grade);
		if (isNaN(gradeNum) || gradeNum < 1 || gradeNum > 10) {
			errorMessage.value = "Оценка должна быть от 1 до 10";
			return false;
		}

		return true;
	};

	const handleSubmit = () => {
		if (!validate()) return;
		showConfirm.value = true;
	};

	const clearForm = () => {
		form.fullName = "";
		form.university = null;
		form.specialty = "";
		form.grade = "";
		form.phone = "";
	};

	const submit = async (emit: any) => {
		errorMessage.value = "";
		isLoading.value = true;

		try {
			await $api("/profile/student", {
				method: "POST",
				body: {
					full_name: form.fullName,
					university_department_id: form.university?.id || null,
					specialty: form.specialty,
					grade: Number(form.grade),
					position: null,
					phone: form.phone || null,
				},
			});
			showConfirm.value = false;
			notify({
				type: "success",
				title: "Успех",
				content: "Ваша заявка успешно отправлена",
			});

			clearForm();
		} catch (e: any) {
			errorMessage.value = e.message || "Ошибка при отправке";
			notify({
				type: "error",
				title: "Ошибка",
				content: e.message || "Ошибка при отправке",
			});
		} finally {
			isLoading.value = false;
		}
	};

	return {
		form,
		isLoading,
		showConfirm,
		errorMessage,
		handleSubmit,
		submit,
	};
}
