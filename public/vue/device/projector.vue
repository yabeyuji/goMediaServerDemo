<template>
	<div>
		<b-row>
			<b-col cols="4"></b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="projectorArrowUp"
				@click="methodSendCommand($event)"
			>
				<icon-circle-up
					data-value="projectorArrowUp"
				/>
			</b-col>
			<b-col cols="4"></b-col>
		</b-row>

		<b-row>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="projectorArrowLeft"
				@click="methodSendCommand($event)"
			>
				<icon-circle-left
					data-value="projectorArrowLeft"
				/>
			</b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="projectorArrowCenter"
				@click="methodSendCommand($event)"
			>
				<icon-circle-center

					data-value="projectorArrowCenter"
				/>
			</b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="projectorArrowRight"
				@click="methodSendCommand($event)"
			>
				<icon-circle-right
					data-value="projectorArrowRight"
				/>
			</b-col>
		</b-row>

		<b-row>
			<b-col cols="4"></b-col>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-value="projectorArrowDown"
				@click="methodSendCommand($event)"
			>
				<icon-circle-down
					data-value="projectorArrowDown"
				/>
			</b-col>
			<b-col cols="4"></b-col>
		</b-row>

		<br>

		<b-row>
			<b-col cols="3"></b-col>

			<b-col
				cols="6"
				class="btn-secondary"
				:class="commonClass"
				data-value="projectorMenu"
				@click="methodSendCommand($event)"
			>
				<icon-menu
					data-value="projectorMenu"
				/>
			</b-col>

			<b-col cols="3"></b-col>
		</b-row>
	</div>
</template>


<script>
module.exports = {
	components: {
		'icon-circle-up': httpVueLoader('public/vue/icon/circle-up.vue'),
		'icon-circle-right': httpVueLoader('public/vue/icon/circle-right.vue'),
		'icon-circle-left': httpVueLoader('public/vue/icon/circle-left.vue'),
		'icon-circle-down': httpVueLoader('public/vue/icon/circle-down.vue'),
		'icon-circle-center': httpVueLoader('public/vue/icon/common-circle.vue'),
		'icon-menu': httpVueLoader('public/vue/icon/menu-main.vue'),
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
			defaultDevice :"projector",
			defaultKey :"status",
		}
	}
}



