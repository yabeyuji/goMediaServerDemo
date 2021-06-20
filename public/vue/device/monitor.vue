<template>
	<div>
		<!-- ============================================= -->
		<b-row>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="monitorPower"
				@click="methodSendCommand($event)"
			>
				<icon-power
					data-value="monitorPower"
				/>
			</b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="monitorVolumeDown"
				@click="methodSendCommand($event)"
			>
				<icon-volume-down
					data-value="monitorVolumeDown"
				/>
			</b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="monitorVolumeUp"
				@click="methodSendCommand($event)"
			>
				<icon-volume-up
					data-value="monitorVolumeUp"
				/>
			</b-col>
		</b-row>

		<br>

		<b-row>
			<b-col cols="4"></b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="monitorArrowUp"
				@click="methodSendCommand($event)"
			>
				<icon-circle-up
					data-value="monitorArrowUp"
				/>
			</b-col>
			<b-col cols="4"></b-col>
		</b-row>

		<b-row>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="monitorArrowLeft"
				@click="methodSendCommand($event)"
			>
				<icon-circle-left
					data-value="monitorArrowLeft"
				/>
			</b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="monitorArrowCenter"
				@click="methodSendCommand($event)"
			>
				<icon-circle-center
					data-value="monitorArrowCenter"
				/>
			</b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="monitorArrowRight"
				@click="methodSendCommand($event)"
			>
				<icon-circle-right
					data-value="monitorArrowRight"
				/>
			</b-col>
		</b-row>


		<b-row>
			<b-col cols="4"></b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="monitorArrowDown"
				@click="methodSendCommand($event)"
			>
				<icon-circle-down
					data-value="monitorArrowDown"
				/>
			</b-col>
			<b-col cols="4"></b-col>
		</b-row>

		<br>

		<b-row>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="monitorSwitch"
				@click="methodSendCommand($event)"
			>
				<span
					data-value="monitorSwitch"
				>SWITCH</span>
			</b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="monitorHDMI"
				@click="methodSendCommand($event)"
			>
				<span
					data-value="monitorHDMI"
				>HDMI</span>
			</b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="monitorUSB"
				@click="methodSendCommand($event)"
			>
				<span
					data-value="monitorUSB"
				>USB</span>
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
		'icon-circle-center': httpVueLoader('public/vue/icon/common-circle.vue'),
		'icon-circle-up': httpVueLoader('public/vue/icon/circle-up.vue'),
		'icon-circle-right': httpVueLoader('public/vue/icon/circle-right.vue'),
		'icon-circle-left': httpVueLoader('public/vue/icon/circle-left.vue'),
		'icon-circle-down': httpVueLoader('public/vue/icon/circle-down.vue'),
	},

	methods: {
		methodSendCommand: function(e) {
			let args = {};
			args['Object'] = e.target.dataset.object == undefined ? this.device : e.target.dataset.object;
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
			defaultDevice :"monitor",
			defaultKey :"status",
		}
	}
}



