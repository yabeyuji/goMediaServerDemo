<template>
	<div>
		<!-- ----------------------------------------- -->

		<b-row>
			<b-col :cols="methodFileColSizeInputBox()" class="btn-success" :class="commonClass">
				<b-form-file
					class="text-left"
					v-model="uploadFile"
					:state="Boolean(uploadFile)"
					accept=".mp4"
					placeholder="Choose a file"
				></b-form-file>
			</b-col>

			<b-col
				:cols="methodFileColSizeButton()"
				class="force-middle"
				:class="[commonClass, methodFileValidate()]"
				@click="methodUpload"
				type="submit"
			>
				<div v-if="isUploading">
					<icon-loading> </icon-loading>
				</div>
				<div v-else-if="!methodUploadFileIsUnique()">
					<icon-not-unique> </icon-not-unique>
				</div>
				<div v-else>
					<icon-record > </icon-record>
				</div>


			</b-col>
		</b-row>

		<!-- ----------------------------------------- -->

		<hr />
		<b-row>
			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-object="vlc"
				data-key="status"
				data-value="vlcStop"
				@click="methodStatusChange($event)"
			>
				<icon-stop
					data-object="vlc"
					data-key="status"
					data-value="vlcStop"
				/>
			</b-col>

			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-object="vlc"
				data-key="status"
				data-value="vlcPause"
				@click="methodStatusChange($event)"
			>
				<icon-pause
					data-object="vlc"
					data-key="status"
					data-value="vlcPause"
				/>
			</b-col>

			<b-col
				cols="4"
				class="btn-secondary"
				:class="commonClass"
				data-object="vlc"
				data-key="status"
				data-value="vlcPlay"
				@click="methodStatusChange($event)"
			>
				<icon-play
					data-object="vlc"
					data-key="status"
					data-value="vlcPause"
				/>
			</b-col>
		</b-row>

		<hr />

		<b-row>
			<b-col
				cols="2"
				class="btn-secondary"
				:class="commonClass"
				data-object="vlc"
				data-key="status"
				data-value="vlcPlayPrevious"
				@click="methodStatusChange($event)"
			>
				<icon-backward
					data-object="vlc"
					data-key="status"
					data-value="vlcPlayPrevious"
				/>
			</b-col>

			<b-col cols="8" class="d-flex align-items-center">
				<b-form-input
					id="range-2"
					v-model="computedVlcProgress"
					type="range"
					min="0"
					step="0.01"
					max="1"
					data-object="vlc"
					data-key="status"
					@mousedown="methodPauseOn($event)"
					@mouseup="methodPauseOff($event)"
				></b-form-input>
			</b-col>

			<b-col
				cols="2"
				class="btn-secondary"
				:class="commonClass"
				data-object="vlc"
				data-key="status"
				data-value="vlcPlayNext"
				@click="methodStatusChange($event)"
			>
				<icon-forward
					data-object="vlc"
					data-key="status"
					data-value="vlcPlayNext"
				/>
			</b-col>
		</b-row>

		<hr />

		<b-row>
			<b-col
				cols="6"
				v-for="(file, index) in files"
				:class="[commonClass, methodClassSelect(file.Valid)]"
				:key="index"
			>

				<img
					class="video-thumbnail"
					data-object="file"
					data-key="valid"
					:data-value="file.Vid"
					@contextmenu="methodToggleValid($event)"
					@click="methodPlayMobile($event)"
					:src="animeDir + file.Vid + '.gif'"
				/>
				<div
					data-object="file"
					data-key="valid"
					:data-value="file.Vid"
					@contextmenu="methodToggleValid($event)"
					@click="methodPlayMobile($event)"
				>
					{{ file.Name }}
				</div>
			</b-col>
		</b-row>
	</div>
</template>


<script>
module.exports = {
	components: {
		'icon-record': httpVueLoader('public/vue/icon/media-record.vue'),
		'icon-stop': httpVueLoader('public/vue/icon/common-stop.vue'),
		'icon-pause': httpVueLoader('public/vue/icon/common-pause.vue'),
		'icon-play': httpVueLoader('public/vue/icon/common-play.vue'),
		'icon-backward': httpVueLoader('public/vue/icon/operation-backward.vue'),
		'icon-forward': httpVueLoader('public/vue/icon/operation-forward.vue'),
		'icon-loading': httpVueLoader('public/vue/icon/common-loading.vue'),
		'icon-not-unique': httpVueLoader('public/vue/icon/common-not-unique.vue'),
	},

methods: {
	methodPlayMobile: function(e) {
		window.open(this.videoDir + e.target.dataset.value + ".mp4", '_blank');
	},

	methodUpload: function() {
		if(this.isUploading == false && this.methodUploadFileIsUnique() == true) {
			this.$emit('file-upload', this.uploadFile);
		}
	},

	methodFileValidate() {
		return this.uploadFile != null ? 'btn-primary' : 'display-none';
	},

	methodFileColSizeInputBox: function() {
		return this.uploadFile != null ? 10 : 12;
	},

	methodFileColSizeButton: function() {
		return this.uploadFile != null ? 2 : 0;
	},

	methodUploadFileIsUnique: function() {
		if (this.uploadFile != null && this.files != [] ) {
			let uploadFile = this.uploadFile
			let result = true
			this.files.forEach(function(file) {
				if (file.Name + ".mp4" == uploadFile.name) {
					result = false
				}
			});
			return result
		}
		return true
	},

	methodStatusChange: function (e) {
		let args = {};
		args['Object'] = e.target.dataset.object;
		args['Key'] = e.target.dataset.key;
		args['Value'] = e.target.dataset.value;
		this.$emit('send-command', args);
	},

	methodPauseOn: function (e) {
		let args = {};
		args['Object'] = e.target.dataset.object;
		args['Key'] = "status";
		args['Value'] = "vlcPauseOn";
		this.$emit('send-command', args);
	},

	methodPauseOff: function (e) {
		let args = {};
		args['Object'] = e.target.dataset.object;
		args['Key'] = "status";
		args['Value'] = "vlcPauseOff";
		this.$emit('send-command', args);
	},

	methodToggleValid: function (e) {
		e.preventDefault();
		let args = {};
		args['Room'] = "common";
		args['Object'] = e.target.dataset.object;
		args['Key'] = e.target.dataset.key;
		args['Value'] = e.target.dataset.value;
		this.$emit('send-command', args);
	},

	methodClassSelect: function(value) {
		return value == true ? 'btn-info' : 'btn-secondary';
	},

	methodSelecteBadge: function(value) {
		if(this.selectedBadge === value){
			this.selectedBadge = '';
		}else{
			this.selectedBadge = value;
		}
	},

}, //end methods

	props: {
		'commonClass': { type: String, default: '' },
		'vlcProgress': { type: Number, default: 0 },
		'isUploading': { type: Boolean, default: false },
		'files': { type: Array, default: [] },
	},

	computed: {
		computedVlcProgress: {
			get() {
				return this.$props.vlcProgress;
			},
			set(value) {
				let args = {};
				args['Object'] = "vlc";
				args['Key'] = "progress";
				args['Value'] = value;
				args['VibrateDisable'] = true;
				this.$emit('send-command', args);
			},
		},
	},

  data: function () {
		return {
			uploadFile: null,
			animeDir: './public/file/anime/',
			videoDir : './public/file/video/',
		};
	},
};

