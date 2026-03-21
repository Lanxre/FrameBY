export enum FramebyAppRole {
  USER = "user",
  STUDENT = "student",
  UNIVERSITY = "university",
  CUSTOMER = "customer",
  BRSM = "brsm",
}

export const ROLE_WEIGHTS: Record<FramebyAppRole, number> = {
	[FramebyAppRole.USER]: 1,
	[FramebyAppRole.STUDENT]: 2,
	[FramebyAppRole.UNIVERSITY]: 3,
	[FramebyAppRole.CUSTOMER]: 4,
	[FramebyAppRole.BRSM]: 5,
};
