<template>
	<div>
		<span style="display:none">--------------------------------</span>
		<b-row>
			<b-col
				cols="4"
				class="btn-secondary "
				:class="commonClass"
				data-value="tvPower"
				@click="methodSendCommand($event)"
			>
				<icon-power
					data-value="tvPower"
				/>
			</b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-object="monitor"
				data-value="monitorVolumeDown"
				@click="methodSendCommand($event)"
			>
				<icon-volume-down
					data-object="monitor"
					data-value="monitorVolumeDown"
				/>
			</b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-object="monitor"
				data-value="monitorVolumeUp"
				@click="methodSendCommand($event)"
			>
				<icon-volume-up
					data-object="monitor"
					data-value="monitorVolumeUp"
				/>
			</b-col>
		</b-row>

		<span style="display:none">--------------------------------</span>

		<b-row>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvReplay"
				@click="methodSendCommand($event)"
			>
				<icon-replay
					data-value="tvReplay"
				/>
			</b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvRecord"
				@click="methodSendCommand($event)"
			>
				<icon-record
					data-value="tvRecord"
				/>
			</b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvChangeMode"
				@click="methodSendCommand($event)"
			>
				<icon-cycle
					data-value="tvChangeMode"
				/>
			</b-col>
		</b-row>

		<span style="display:none">--------------------------------</span>
		<br>

		<b-row>
			<b-col cols="4"></b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvArrowUp"
				@click="methodSendCommand($event)"
			>
				<icon-circle-up
					data-value="tvArrowUp"
				/>
			</b-col>
			<b-col cols="4"></b-col>
		</b-row>

		<span style="display:none">--------------------------------</span>

		<b-row>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvArrowLeft"
				@click="methodSendCommand($event)"
			>
				<icon-circle-left
					data-value="tvArrowLeft"
				/>
			</b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvArrowCenter"
				@click="methodSendCommand($event)"
			>
				<icon-circle
					data-value="tvArrowCenter"
				/>
			</b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvArrowRight"
				@click="methodSendCommand($event)"
			>
				<icon-circle-right
					data-value="tvArrowRight"
				/>
			</b-col>
		</b-row>

		<span style="display:none">--------------------------------</span>

		<b-row>
			<b-col cols="4"></b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvArrowDown"
				@click="methodSendCommand($event)"
			>
				<icon-circle-down
					data-value="tvArrowDown"
				/>
			</b-col>
			<b-col cols="4"></b-col>
		</b-row>

		<span style="display:none">--------------------------------</span>
		<br>

		<b-row>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvMenu"
				@click="methodSendCommand($event)"
			>
				<icon-menu
					data-value="tvMenu"
				/>
			</b-col>

			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvMenuDub"
				@click="methodSendCommand($event)"
			>
				<icon-menu-sub
					data-value="tvMenuDub"
				/>
			</b-col>

			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvReturn"
				@click="methodSendCommand($event)"
			>
				<icon-return
					data-value="tvReturn"
				/>
			</b-col>

		</b-row>


		<span style="display:none">--------------------------------</span>
		<br>

		<b-row>
			<b-col
				cols="3"
				class="btn-primary"
				:class="commonClass"
				data-value="tvColorBlue"
				@click="methodSendCommand($event)"
			>
				<span
					data-value="tvColorBlue"
				>青</span>
			</b-col>

			<b-col
				cols="3"
				class="btn-danger"
				:class="commonClass"
				data-value="tvColorRed"
				@click="methodSendCommand($event)"
			>
				<span
					data-value="tvColorRed"
				>赤</span>
			</b-col>

			<b-col
				cols="3"
				class="btn-success"
				:class="commonClass"
				data-value="tvColorGreen"
				@click="methodSendCommand($event)"
			>
				<span
					data-value="tvColorGreen"
				>緑</span>
			</b-col>

			<b-col
				cols="3"
				class="btn-warning"
				:class="commonClass"
				data-value="tvColorYellow"
				@click="methodSendCommand($event)"
			>
				<span
					data-value="tvColorYellow"
				>黃</span>
			</b-col>

		</b-row>

		<span style="display:none">--------------------------------</span>
		<br>

		<b-row>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvOperateStop"
				@click="methodSendCommand($event)"
			>
				<icon-stop
					data-value="tvOperateStop"
				/>
			</b-col>

			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvOperatePause"
				@click="methodSendCommand($event)"
			>
				<icon-pause
					data-value="tvOperatePause"
				/>
			</b-col>

			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvOperatePlay"
				@click="methodSendCommand($event)"
			>
				<icon-play
					data-value="tvOperatePlay"
				/>
			</b-col>

		</b-row>

		<span style="display:none">--------------------------------</span>
		<br>

		<b-row>
			<b-col
				cols="3"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvOperateFBackward"
				@click="methodSendCommand($event)"
			>
				<icon-f-backward
					data-value="tvOperateFBackward"
				/>
			</b-col>

			<b-col
				cols="3"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvOperateBackward"
				@click="methodSendCommand($event)"
			>
				<icon-backward
					data-value="tvOperateBackward"
				/>
			</b-col>

			<b-col
				cols="3"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvOperateForward"
				@click="methodSendCommand($event)"
			>
				<icon-forward
					data-value="tvOperateForward"
				/>
			</b-col>

			<b-col
				cols="3"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvOperateFForward"
				@click="methodSendCommand($event)"
			>
				<icon-f-forward
					data-value="tvOperateFForward"
				/>
			</b-col>

		</b-row>

		<span style="display:none">--------------------------------</span>
		<br>

		<b-row>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvChannelUp"
				@click="methodSendCommand($event)"
			>
				<icon-channel-up
					data-value="tvChannelUp"
				/>
			</b-col>

			<b-col
				cols="2"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvChannel01NHKG"
				@click="methodSendCommand($event)"
			>
				<img
					src="../public/image/tv/nhk_g.png"
					class="channel-thumbnail"
					data-value="tvChannel01NHKG"
				/>
			</b-col>

			<b-col
				cols="2"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvChannel02NHKE"
				@click="methodSendCommand($event)"
			>
				<img
					src="../public/image/tv/nhk_e.png"
					class="channel-thumbnail"
					data-value="tvChannel02NHKE"
				/>
			</b-col>

			<b-col
				cols="2"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvChannel04NihonTV"
				@click="methodSendCommand($event)"
			>
				<img
					src="../public/image/tv/nihon_tv.png"
					class="channel-thumbnail"
					data-value="tvChannel04NihonTV"
				/>
			</b-col>

			<b-col
				cols="2"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvChannel05TVAsahi"
				@click="methodSendCommand($event)"
			>
				<img
					src="../public/image/tv/tv_asahi.png"
					class="channel-thumbnail"
					data-value="tvChannel05TVAsahi"
				/>
			</b-col>

		</b-row>

		<b-row>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvChannelDown"
				@click="methodSendCommand($event)"
			>
				<icon-channel-down
					data-value="tvChannelDown"
				/>
			</b-col>

			<b-col
				cols="2"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvChannel07TVTokyo"
				@click="methodSendCommand($event)"
			>
				<img
					src="../public/image/tv/tv_tokyo.png"
					class="channel-thumbnail"
					data-value="tvChannel07TVTokyo"
				/>
			</b-col>

			<b-col
				cols="2"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvChannel08FujiTV"
				@click="methodSendCommand($event)"
			>
				<img
					src="../public/image/tv/fuji_tv.png"
					class="channel-thumbnail"
					data-value="tvChannel08FujiTV"
				/>
			</b-col>

			<b-col
				cols="2"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvChannel06TBS"
				@click="methodSendCommand($event)"
			>
				<img
					src="../public/image/tv/tbs.jpeg"
					class="channel-thumbnail"
					data-value="tvChannel06TBS"
				/>
			</b-col>

			<b-col
				cols="2"
				class="btn-secondary"
				:class="commonClass"
				data-value="tvChannel09TokyoMX"
				@click="methodSendCommand($event)"
			>
				<img
					src="../public/image/tv/tokyo_mx.jpeg"
					class="channel-thumbnail"
					data-value="tvChannel09TokyoMX"
				/>
			</b-col>
		</b-row>

	</div>
</template>


<script>
module.exports = {
	components: {
		'icon-power': httpVueLoader('public/vue/icon/common-powerOnOff.vue'),
		'icon-volume-up': httpVueLoader('public/vue/icon/volume-up.vue'),
		'icon-volume-down': httpVueLoader('public/vue/icon/volume-down.vue'),

		'icon-circle': httpVueLoader('public/vue/icon/common-circle.vue'),
		'icon-circle-up': httpVueLoader('public/vue/icon/circle-up.vue'),
		'icon-circle-right': httpVueLoader('public/vue/icon/circle-right.vue'),
		'icon-circle-left': httpVueLoader('public/vue/icon/circle-left.vue'),
		'icon-circle-down': httpVueLoader('public/vue/icon/circle-down.vue'),
		'icon-menu': httpVueLoader('public/vue/icon/menu-main.vue'),
		'icon-menu-sub': httpVueLoader('public/vue/icon/memu-sub.vue'),
		'icon-replay': httpVueLoader('public/vue/icon/media-replay.vue'),
		'icon-record': httpVueLoader('public/vue/icon/media-record.vue'),
		'icon-cycle': httpVueLoader('public/vue/icon/common-arrow-cycle.vue'),
		'icon-return': httpVueLoader('public/vue/icon/common-return.vue'),

		'icon-stop': httpVueLoader('public/vue/icon/common-stop.vue'),
		'icon-pause': httpVueLoader('public/vue/icon/common-pause.vue'),
		'icon-play': httpVueLoader('public/vue/icon/common-play.vue'),

		'icon-f-backward': httpVueLoader('public/vue/icon/operation-f-backward.vue'),
		'icon-backward': httpVueLoader('public/vue/icon/operation-backward.vue'),
		'icon-forward': httpVueLoader('public/vue/icon/operation-forward.vue'),
		'icon-f-forward': httpVueLoader('public/vue/icon/operation-f-forward.vue'),

		'icon-channel-up': httpVueLoader('public/vue/icon/chevron-up.vue'),
		'icon-channel-down': httpVueLoader('public/vue/icon/chevron-down.vue'),

},

	methods: {
		methodSendCommand: function(e) {
			let args = {};
			args['Object'] = e.target.dataset.object == undefined ? this.defaultDevice : e.target.dataset.object;
			args['Key'] = e.target.dataset.key == undefined ? this.defaultKey : e.target.dataset.key;
			args['Value'] = e.target.dataset.value;
			this.$emit('send-command', args);
		},
	},

	props: {
		'commonClass': {type: String, default: ''},
	},

	data: function() {
		return {
			defaultDevice :"tv",
			defaultKey :"status",
		}
	}
}



