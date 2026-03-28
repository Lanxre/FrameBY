import { reactive, ref } from "vue";
import { $api } from "../useApi";

export function useUniversityForm() {
	const form = reactive({
		fullName: "",
		university: null as any,
		department: "",
		position: "",
		phone: "",
	});

	const isLoading = ref(false);
	const showConfirm = ref(false);
	const errorMessage = ref("");

	const { notify } = useNotificationStore();

	const validate = () => {
		if (
			!form.fullName ||
			!form.university ||
			!form.department ||
			!form.position ||
			!form.phone
		) {
			errorMessage.value = "Заполните все поля";
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
		form.department = "";
		form.position = "";
		form.phone = "";
	};

	const submit = async (emit: any) => {
		errorMessage.value = "";
		isLoading.value = true;

		try {
			await $api("/profile/university", {
				method: "POST",
				body: {
					full_name: form.fullName,
					department: form.department,
					position: form.position || null,
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
			errorMessage.value = e.message || "Ошибка при сохранении";
			notify({
				type: "error",
				title: "Ошибка",
				content: e.message || "Ошибка при сохранении",
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
		submit,
		handleSubmit,
	};
}
