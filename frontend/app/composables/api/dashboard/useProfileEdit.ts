import { $api } from "@/composables/api/useApi";
import { useUniversityDepartments } from "@/composables/api/useUniversityDepartments";
import { useSpecialties } from "@/composables/api/useSpecialties";
import { PROFILE_TYPES, type ProfileRow } from "@/types/dashboard/profile";
import { FramebyAppRole } from "~/types/frontend/enums/role";

export interface RoleOption {
	id: string;
	name: string;
}

export function useProfileEdit() {
	const roleOptions: RoleOption[] = PROFILE_TYPES.filter(
		(t) => t.value !== "all",
	).map((t) => ({
		id: t.value,
		name: t.label,
	}));

	const { departments, fetchDepartments } = useUniversityDepartments();
	const { specialties, fetchByUniversityDepartment } = useSpecialties();

	const selectedRole = ref<RoleOption | null>(null);
	const selectedUniversity = ref<{ id: string; name: string } | null>(null);
	const fullName = ref("");
	const subrole = ref("");
	const specialtyId = ref<string | null>(null);
	const specialtyName = ref("");
	const grade = ref("");
	const position = ref("");
	const phone = ref("");
	const isLoading = ref(false);
	const errorMessage = ref("");

	const getProfileField = (profile: ProfileRow["profile"], field: string) => {
		return (profile as any)?.[field] ?? null;
	};

	const populateForm = (profile: ProfileRow) => {
		const roleOption = roleOptions.find((r) => r.id === profile.role);
		selectedRole.value = roleOption || null;

		const universityInfo = getProfileField(profile.profile, "university");
		if (universityInfo?.id) {
			const dept = departments.value.find((d) => d.id === universityInfo.id);
			selectedUniversity.value = dept || {
				id: universityInfo.id,
				name: universityInfo.university_name,
			};
		} else {
			selectedUniversity.value = null;
		}

		fullName.value = getProfileField(profile.profile, "full_name") || "";
		subrole.value = getProfileField(profile.profile, "subrole") || "";
		specialtyId.value = getProfileField(profile.profile, "specialty_id") || null;
		specialtyName.value = getProfileField(profile.profile, "specialty") || "";
		grade.value = getProfileField(profile.profile, "grade")?.toString() || "";
		position.value = getProfileField(profile.profile, "position") || "";
		phone.value = getProfileField(profile.profile, "phone") || "";
	};

	const resetForm = () => {
		errorMessage.value = "";
	};

	const buildUpdateBody = (): Record<string, any> | null => {
		if (!selectedRole.value) return null;

		const body: any = {
			role: selectedRole.value.id,
		};

		if (selectedRole.value.id !== FramebyAppRole.USER) {
			body.full_name = fullName.value || null;
		}

		if (
			selectedRole.value.id === FramebyAppRole.BRSM ||
			selectedRole.value.id === FramebyAppRole.UNIVERSITY ||
			selectedRole.value.id === FramebyAppRole.CUSTOMER
		) {
			if (subrole.value) body.subrole = subrole.value;
		}

		if (selectedRole.value.id === FramebyAppRole.STUDENT) {
			if (selectedUniversity.value?.id) {
				body.university_department_id = selectedUniversity.value.id;
			}
			if (specialtyId.value) body.specialty_id = specialtyId.value;
			if (grade.value) body.grade = parseFloat(grade.value);
		}

		if (selectedRole.value.id === FramebyAppRole.UNIVERSITY) {
			if (selectedUniversity.value?.id) {
				body.university_department_id = selectedUniversity.value.id;
			}
		}

		if (selectedRole.value.id !== FramebyAppRole.USER && position.value) {
			body.position = position.value;
		}

		if (selectedRole.value.id !== FramebyAppRole.USER && phone.value) {
			body.phone = phone.value;
		}

		return body;
	};

	const save = async (userId: string): Promise<boolean> => {
		if (!selectedRole.value) return false;

		isLoading.value = true;
		errorMessage.value = "";

		try {
			const body = buildUpdateBody();
			if (!body) return false;

			await $api(`/admin/profiles/${userId}/subrole`, {
				method: "PATCH",
				body,
			});
			return true;
		} catch (e: any) {
			errorMessage.value = e.data?.error || e.message || "Ошибка сохранения";
			return false;
		} finally {
			isLoading.value = false;
		}
	};

	onMounted(() => {
		fetchDepartments();
	});

	return {
		roleOptions,
		departments,
		specialties,
		selectedRole,
		selectedUniversity,
		fullName,
		subrole,
		specialtyId,
		specialtyName,
		grade,
		position,
		phone,
		isLoading,
		errorMessage,
		populateForm,
		resetForm,
		save,
		fetchByUniversityDepartment,
	};
}
