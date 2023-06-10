// GET stamp response body 型定義
// POST stamp request body 型定義
type Record = {
	time: string;
	date: string;
	type: string;
};
type Checkpoint = {
	lon: float;
	lat: float;
};
type PostStampBody = {
	record: Record;
	group_id: int;
	place_id: int;
	note: string;
	checkpoint: Checkpoint;
};

// stamp endpoint path param
type date = string;

// stamp methods
const AditMethodEnum = {
	manager: 'manager',
	card: 'card',
	finger: 'finger',
	web: 'web',
	mobile: 'mobile',
	mobile_fix: 'mobile_fix',
	pc: 'pc',
	pc_fix: 'pc_fix',
	api: 'api',
	line: 'line',
	line_beacon: 'line_beacon',
	slack: 'slack',
	face: 'face',
	senselink: 'senselink',
	geofencing: 'geofencing',
	card_windows: 'card_windows',
	card_pittouch: 'card_pittouch',
	card_ios: 'card_ios',
	card_ios_nfc: 'card_ios_nfc',
	card_android: 'card_android',
	card_other: 'card_other',
	finger_windows: 'finger_windows',
	face_ios: 'face_ios',
	face_other: 'face_other',
	employee_select_ios: 'employee_select_ios',
	employee_select_other: 'employee_select_other'
} as const;

const AditTypeEnum = {
	auto: 'auto',
	manual: 'manual',
	gps: 'gps',
	checkpoint: 'checkpoint',
	auto_start: 'auto_start',
	auto_end: 'auto_end'
} as const;

export function getPostStampBody(
	time: string,
	date: string,
	type: string,
	groupId: int,
	placeId: int,
	note: string,
	lon: number,
	lat: number
): PostStampBody {
	return (PostStampBody = {
		record: {
			time: time,
			date: date,
			type: type
		},
		group_id: groupId,
		place_id: placeId,
		note: note,
		checkpoint: {
			lon: lon,
			lat: lat
		}
	});
}

export function getStampUrl(
	client_id: int,
	employee_id: int,
	dateOrMethod: date | AditMethodEnum
): string {
	return `http://localhost:8088/stamp/v1/clients/${client_id}/employees/${employee_id}/stamp/${dateOrMethod}`;
}
