<script lang="ts">
	// 打刻API
	import { getPostStampBody } from 'src/api/stamp/model';
	import { getStampUrl } from 'src/api/stamp/model';
	import { postStamp } from 'src/api/stamp/stamp';
	// 環境変数
	import { env } from '../util/env';
	// 現在時刻取得関数
	import { getNow } from '../util/time';

	let result: any = null;
	let clientId: number = 50000;
	let employeeId: number = 1;
	let method: string = 'api';
	let date: string = getNow();
	let time: string = new Date().getTime().toString();
	let type: string = 'auto';
	let groupId: number = 11;
	let placeId: number = 1;
	let note: string = 'test';
	let lon: number = 139.767125;
	let lat: number = 35.681236;
</script>

<div>
	<div>
		<p>現在時刻は{getNow()}です</p>
		<p>環境変数は{env.STAMP_API_ENDPOINT}です</p>
	</div>
</div>
